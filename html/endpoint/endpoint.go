package main

import (
	"html/template"
	"net/http"
)

func redTeaPotHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./views/product.html")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, "TEST")
	// handle error
}

func main() {
	http.HandleFunc("/red-tea-pot", redTeaPotHandler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
