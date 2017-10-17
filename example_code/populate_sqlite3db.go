package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func populate_sql() {
	database, _ := sql.Open("sqlite3", "./data/students.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS students (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO students (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Brandon", "Stewart")
	rows, _ := database.Query("SELECT id, firstname, lastname FROM students")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
