package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func createTables(){
	createRemindersTableSQL := `CREATE TABLE IF NOT EXISTS reminders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description VARCHAR(255) NOT NULL,
		due date NOT NULL,
		tags VARCHAR(255)
	);`
	_, err := DB.Exec(createRemindersTableSQL)
	if err != nil {
		panic(err)
	}
}

func createDB(dbPath string){
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		file, err := os.Create(dbPath)
		if err != nil {
			panic("Could not create database")
		}
		file.Close()
	}
}

func Init(){
	dbPath := "./data/db.db"

	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil { 
		panic(err)
	}
	defer DB.Close()

	createDB(dbPath)
	createTables()
	
}
