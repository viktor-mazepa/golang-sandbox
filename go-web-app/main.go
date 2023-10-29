package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"webApp/rps"
)

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)
	out, err := json.MarshalIndent(result, "", "     ")

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	template, err := template.ParseFiles(page)

	if err != nil {
		log.Println(err)
		return
	}

	err = template.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
