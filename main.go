package main

import (
	"log"
	"net/http"

	"github.com/YabseraBogale/erp/actors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	err = db.AutoMigrate(&actors.EmployeeContact{}, &actors.Employee{},
		&actors.Item{}, &actors.TransactionType{},
		&actors.ItemLog{}, &actors.CheckOut{}, &actors.CheckIn{})
	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("/employee_registeration", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "post" {

		}

	})

	http.HandleFunc("/item_registeration", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "post" {

		}
	})

	http.HandleFunc("/checkin", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "post" {

		}
	})

	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "post" {

		}
	})

	err = http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Println(err)
	}
}
