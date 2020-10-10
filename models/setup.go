package models

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	// Need to use postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB holds a reference to the database
var DB *gorm.DB

// ConnectDataBase connects to postgres db
func ConnectDataBase() {
	// database, err := gorm.Open("sqlite3", "books.db")
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=pguser dbname=books password=pgpass sslmode=disable")
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}

// BulkBookLoad loads data from csv
func BulkBookLoad(filename string) {
	var book Book

	cf, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open csv file.", err)
	}

	r := csv.NewReader(cf)
	// skip csv header row
	_, err = r.Read()

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		book = Book{Title: rec[9], Author: rec[7]}
		DB.Create(&book)
	}
}
