package main

import (
	"fmt"
	"log"
	"net/http"

	handle "github.com/Metzark/cfb/api/handlers"
	postgres_client "github.com/Metzark/cfb/api/pgc"
)


func main(){

    // Create a postgres client (that has a pool)
    pgc := postgres_client.CreatePGC()

    // CSS
    fs := http.FileServer(http.Dir("web/static"))
    http.Handle("GET /static/", http.StripPrefix("/static/", fs))

    // JSON
    http.HandleFunc("GET /search-teams", handle.SearchTeams)

    // HTML
    http.HandleFunc("GET /teams/", func(w http.ResponseWriter, r *http.Request) {
        handle.Teams(w, r, pgc)
    })
    http.HandleFunc("GET /predict", handle.Predict)

    fmt.Println("Running on 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
