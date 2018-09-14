package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "oscar.beaumont"
	password = ""
	dbname   = "oscar.beaumont"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	/*queryStmt, err := db.Prepare(`select * from "users";`) // WHERE id=$1
	if err != nil {
		log.Fatal(err)
	}

	//var name string
	err = queryStmt.QueryRow().Scan("1") //&name)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}*/

	/*rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()*/

	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	var email string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, 2)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}

	/*sqlStatement := `
	  INSERT INTO users (id, username, email, first_name, last_name)
	  VALUES ($1, $2, $3, $4, $5)`
		_, err = db.Exec(sqlStatement, 1, "oscar", "oscar@otbeaumont.me", "Oscar", "Beaumont")
		if err != nil {
			panic(err)
		}*/

}
