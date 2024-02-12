package main

import (
	"log"
	"net/http"
	"html/template"
	"fmt"
)

var templates *template.Template

type usr struct {
	Name string
	Email string
}

func home(w http.ResponseWriter, r *http.Request){
	// r - request
	// w - response
	// w.Write([]byte("Hello World!"))
	u := usr{
		"Ana",
		"ana.gm@gmail.com",
	}

	templates.ExecuteTemplate(w, "home.html", u)
}

func user(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Users"))
}

func main() {
	// html pages
	templates = template.Must(template.ParseGlob("*.html"))
	
	// routes
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", user)

	fmt.Println("Listen port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}