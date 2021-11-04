package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnect() sql.DB {

	db, err := sql.Open("mysql", "farmer:farmer@tcp(localhost)/ferme_bd")

	if err != nil {
		panic(err.Error())
	}

	return *db

}

func CloseDatabase(db *sql.DB) {
	defer db.Close()
}

func Print() {
	db := GetConnect()
	results, err := db.Query("SELECT fname FROM bank.customer")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag string

		err = results.Scan(&tag)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println(tag)

	}

	CloseDatabase(&db)
}
