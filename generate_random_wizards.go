// main is a standalone package,
// it is used to generate and print
// random wizards names
// its functions (except 'main', of course)
// should be used to "create Magic Life"
// -> well, for the MAGIC CRUD...
package main

import (
	"encoding/json"
	"fmt"
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

// Wizard is the content for the Magic Inventory DB
type Wizard struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       float64 `json:"age"`
	Category  string  `json:"category"` // Families, Guests, Villains
	Arrested  bool    `json:"arrested"`
	Dead      bool    `json:"dead"`
}

// Any checks if id of a new Wizard doesn't already exist
func Any(id string, wizards []Wizard) bool {
	for _, w := range wizards {
		if w.ID == id {
			return true
		}
	}
	return false
}

// AddWizard adds a new one only if its generated id
// does not already belong to another created Wizard
func AddWizard(categories []string, name Name, wizards []Wizard) []Wizard {
	wizard := Wizard{uuid.Must(uuid.NewV4()).String(), //GenerateRandomID(32),
		strings.Title(name.First),
		strings.Title(name.Last),
		float64(rand.Int()%20 + 20),
		categories[rand.Int()%3],
		false,
		false}

	if Any(wizard.ID, wizards) {
		return AddWizard(categories, name, wizards)
	}

	return append(wizards, wizard)
}

// GenerateWizards is the function that will generate
// random wizard identities, from GetRandomNames
func GenerateWizards(body []byte) (wizards []Wizard, err error) {
	names := Names{}
	categories := []string{"Families", "Guests", "Villains"}
	err = json.Unmarshal(body, &names)

	if err != nil {
		return wizards, err
	}

	for _, name := range names.Results {
		wizards = AddWizard(categories, name.Name, wizards)
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

	return body, nil
}

func main() {
	body, _ := GetRandomNames(10)
	wizards, _ := GenerateWizards(body)

	js, _ := json.Marshal(wizards)
	fmt.Println(string(js))

	/*
		for _, w := range wizards {
			j, _ := json.Marshal(w)
			fmt.Println(string(j))
		}
	*/
}
