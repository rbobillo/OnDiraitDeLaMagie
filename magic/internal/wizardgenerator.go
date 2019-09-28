// Package internal groups some useful
// internal libraries
// here, the wizardgenerator uses randomuser.me
// to generator a bunch of random Wizard
package internal

import (
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/dao"
	"github.com/satori/go.uuid" // go get github.com/satori/go.uuid
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

// Name is used for RandomNames JSON parsing
// it is one of the Identity components
// { "title" : "", "first" : "", "last" }
// excluding the "title" key
type Name struct {
	First string
	Last  string
}

// Identity is used for RandomNames JSON parsing
// it is the type contains in "results"
// { "name" : { ... } }
type Identity struct {
	Name Name
}

// Names is used for RandomNames JSON parsing
// it is the whole JSON content for generated names
// { "result" : [ { "name" : { ... } }, ... ], "info" : { ... } }
// excluding the "info" key
// cf: https://randomuser.me/documentation#howto
type Names struct {
	Results []Identity
}

// any checks if id of a new Wizard doesn't already exist
func any(id uuid.UUID, wizards []dao.Wizard) bool {
	for _, w := range wizards {
		if w.ID == id {
			return true
		}
	}
	return false
}

// addWizard adds a new one only if its generated id
// does not already belong to another created Wizard
func addWizard(categories []string, name Name, wizards []dao.Wizard) []dao.Wizard {
	wizard := dao.Wizard{ID: uuid.Must(uuid.NewV4()),
		FirstName: strings.Title(name.First),
		LastName:  strings.Title(name.Last),
		Age:       float64(9), //rand.Int()%20 + 20
		Category:  categories[rand.Int()%3],
		Arrested:  false,
		Dead:      false}

	if any(wizard.ID, wizards) {
		return addWizard(categories, name, wizards)
	}

	return append(wizards, wizard)
}

// GenerateWizards is the function that will generate
// random wizard identities, from GetRandomNames
func GenerateWizards(body []byte) (wizards []dao.Wizard, err error) {
	names := Names{}
	categories := []string{"Families", "Guests", "Villains"}
	err = json.Unmarshal(body, &names)

	if err != nil {
		return wizards, err
	}

	for _, name := range names.Results {
		wizards = addWizard(categories, name.Name, wizards)
	}

	return wizards, err
}

// GetRandomNames build a custom URL
// (target: random names generator API)
// and print JSON data resulting from GET request
func GetRandomNames(qty int) (body []byte, errs []error) {
	wizardsQty := strconv.Itoa(qty)
	wizardsOrigin := "gb" // "gb,fr,us" ...
	format := "pretty"

	url := "https://randomuser.me/api/" +
		"?results=" + wizardsQty +
		"&nat=" + wizardsOrigin +
		"&format=" + format +
		"&exc=gender,location,email,login,registered,dob,phone,cell,picture,nat,id"

	resp, err := http.Get(url)

	if err != nil {
		return body, append(errs, err)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return body, append(errs, err)
	}

	Debug("random names generated")

	return body, nil
}
