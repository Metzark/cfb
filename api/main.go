package main

import (
	"fmt"
	"log"
	"net/http"

	handle "github.com/Metzark/cfb/api/handlers"
)


func main(){

    // CSS
    fs := http.FileServer(http.Dir("web/static"))
    http.Handle("GET /static/", http.StripPrefix("/static/", fs))

    // JSON
    http.HandleFunc("GET /search-teams", handle.SearchTeams)

    // HTML
    http.HandleFunc("GET /teams/", handle.Teams)
    http.HandleFunc("GET /predict", handle.Predict)

    fmt.Println("Running on 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

