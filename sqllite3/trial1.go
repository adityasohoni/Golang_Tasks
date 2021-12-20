package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func main() {
	database, _ := sql.Open("sqlite3", "./dbfile.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INT PRIMARY KEY , firstname TEXT, lastname TEXT)")
	statement.Exec()

	statement2, _ := database.Prepare("INSERT INTO people values (2,'Revati','Sohoni')")
	statement2.Exec()

	rows, _ := database.Query("SELECT id, firstname, lastname from people")
	var fn string
	var ln string
	var id int
	for rows.Next() {
		rows.Scan(&id, &fn, &ln)
		fmt.Println(strconv.Itoa(id) + ": " + fn + " " + ln)
	}

}
