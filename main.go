package main

import (
	"html/template"
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
		&actors.Item{}, &actors.CheckOut{}, &actors.CheckIn{})

	if err != nil {
		log.Println(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/employee_registeration", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("public/employee_registeration.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if r.Method == "post" {

		}

		templ.Execute(w, nil)
	})

	http.HandleFunc("/item_registeration", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("public/item_registeration.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if r.Method == "post" {

		}

		templ.Execute(w, nil)
	})

	http.HandleFunc("/checkin", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("public/checkin.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if r.Method == "post" {

		}

		templ.Execute(w, nil)
	})

	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("public/checkout.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if r.Method == "post" {

		}

		templ.Execute(w, nil)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("public/dashboard.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if r.Method == "post" {

		}

		templ.Execute(w, nil)
	})

	err = http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Println(err)
	}
}
