package db

import (
	"os"
)

func Init(){
	DB_PATH := "./data/db.db"
	_, err := os.Stat(DB_PATH)
	if os.IsNotExist(err){
		file, err := os.Create(DB_PATH)
		if err != nil {
			panic("Could not create database")
		}
		file.Close()
	}
}
