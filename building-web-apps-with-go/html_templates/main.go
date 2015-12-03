//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"html/template"
	"net/http"
	"path"
)

type Book struct {
	Title	string
	Author	string
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":8080", nil)
}

func ShowBooks(rw http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(rw, book); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
