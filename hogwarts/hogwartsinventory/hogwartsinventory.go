// Package hogwartsinventory is used to setup and manipulate
// the magic database (hogwartsinventory)
package hogwartsinventory

import "database/sql"

func InitHogwartsInventory(psqlInfo string) (*sql.DB, err){
	db, err := sql.Open("postgres", psqlInfo)

}