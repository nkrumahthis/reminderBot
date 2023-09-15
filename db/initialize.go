package db

import (
	"fmt"
	"os"
)

func Init(){
	DB_PATH := "./data/db.db"

	fmt.Println(DB_PATH)
	_, err := os.Stat(DB_PATH)
	fmt.Println(err)
	if os.IsNotExist(err){
		file, err := os.Create(DB_PATH)
		if err != nil {
			panic("Could not create database")
		}
		file.Close()
	}
}
