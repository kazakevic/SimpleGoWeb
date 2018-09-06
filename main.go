package main

import (
	"encoding/csv"
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
	url := "https://raw.githubusercontent.com/vilnius/mokyklos/master/data/Mokyklu%20planuojamos%20komplektacijos2018.csv"
	data, _ := readCSVFromUrl(url)

	t.ExecuteTemplate(w, "index", data)
}

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)

	//if css and js in assets dir not from cdn
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	log.Fatal(http.ListenAndServe(":8080", router))
}
