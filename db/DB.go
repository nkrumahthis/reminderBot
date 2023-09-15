package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func createTables(){
	create_reminders_table := `CREATE TABLE IF NOT EXISTS reminders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description VARCHAR(255) NOT NULL,
		due date NOT NULL,
		tags VARCHAR(255)
	);`
	_, err := DB.Exec(create_reminders_table)
	if err != nil {
		panic(err)
	}
}

func createDB(DB_PATH string){
	_, err := os.Stat(DB_PATH)
	if os.IsNotExist(err){
		file, err := os.Create(DB_PATH)
		if err != nil {
			panic("Could not create database")
		}
		file.Close()
	}
}

func Init(){
	DB_PATH := "./data/db.db"

	var err error
	DB, err = sql.Open("sqlite3", DB_PATH)
	if err != nil { 
		panic(err)
	}
	// defer DB.Close()

	createDB(DB_PATH)
	createTables()
	
}
