package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	t, err := template.ParseFiles("views/index.html", "views/header.html", "views/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)

	//if css and js in assets dir not from cdn
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	log.Fatal(http.ListenAndServe(":8080", router))
}
