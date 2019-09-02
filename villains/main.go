package main

import (
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/internal"
	"io/ioutil"

	"net/http"
	"os"
)


func init() {

}
func main() {
	//hostname := internal.GetEnvOrElse("PG_HOST", "localhost")
	//portaddr := internal.GetEnvOrElse("PG_PORT", "5432")
	//username := internal.GetEnvOrElse("POSTGRES_USER", "magic")
	//password := internal.GetEnvOrElse("POSTGRES_PASSWORD", "magic")
	//database := internal.GetEnvOrElse("POSTGRES_DB", "magicinventory")
	//
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	hostname, portaddr, username, password, database)
	//
	//db, err := magicinventory.InitMagicInventory(psqlInfo)
	//
	//if err != nil {
	//	internal.Error(err.Error())
	//}
	//
	//err = api.InitMagic(db)
	//
	//defer db.Close()
	var wizards []dao.Wizard
	url := "http://localhost:9090/wizards"
	fmt.Println("URL:>", url)
	//var netClient = &http.Client{}
	//tr := &http.Transport{
	//	MaxIdleConns:       20,
	//	MaxIdleConnsPerHost:  20,
	//}
	//netClient = &http.Client{Transport: tr}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &wizards)

	//decoder := json.NewDecoder(resp.Body)
	//err = decoder.Decode(&wizards)

	if err != nil {
		internal.Warn("cannot convert Body to JSON")
		fmt.Println(err)
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	for p := range wizards {
		fmt.Printf("Name = %v", wizards[p].FirstName)
		fmt.Println()
	}

	//log.Fatal(http.ListenAndServe(":909", nil))
}
