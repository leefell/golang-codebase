package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {

	u := usuario{
		"Alexandre",
		"alexandre.augusto@gmail.com",
	}

	templates.ExecuteTemplate(w, "home.html", u)
}

type usuario struct {
	Nome  string
	Email string
}

func main() {
	// Essa expressão referencia todos os arquivos .html
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
