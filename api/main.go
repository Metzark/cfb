package main

import (
	"fmt"
	"log"
	"net/http"

	handle "github.com/Metzark/cfb/api/handlers"
	pg "github.com/Metzark/cfb/api/pg"
)


func main(){
    // Create a postgres client (that has a pool)
    pgc := pg.CreatePGC()

    // CSS
    fs := http.FileServer(http.Dir("web/static"))
    http.Handle("GET /static/", http.StripPrefix("/static/", fs))

    // JSON
    http.HandleFunc("GET /search-teams", func(w http.ResponseWriter, r *http.Request) {
        handle.SearchTeams(w, r, pgc)
    })

    http.HandleFunc("POST /query", func(w http.ResponseWriter, r *http.Request) {
        handle.CustomQuery(w, r, pgc)
    })

    // HTML
    http.HandleFunc("GET /teams/", func(w http.ResponseWriter, r *http.Request) {
        handle.Teams(w, r, pgc)
    })

    http.HandleFunc("GET /predict", handle.Predict)

    http.HandleFunc("GET /query", func(w http.ResponseWriter, r *http.Request) {
        handle.Query(w, r, pgc)
    })


    fmt.Println("Running on 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
