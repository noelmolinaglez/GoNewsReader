package main

import (
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	err := godotenv.Load()
	port := os.Getenv("PORT")
	if err != nil {
		log.Println("Error loading .env file")
	}

	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))



	mux := http.NewServeMux()
    mux.Handle("/assets/",http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

}
