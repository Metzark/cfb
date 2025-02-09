package main

import (
	"fmt"
	"log"
	"net/http"

	handle "github.com/Metzark/cfb/api/handlers"
	pg "github.com/Metzark/cfb/api/pg"
)


func main(){

<<<<<<< HEAD
=======
    // Create a postgres client (that has a pool)
    pgc := pg.CreatePGC()

>>>>>>> 5fc5e57 (Update search teams to pull from postgres)
    // CSS
    fs := http.FileServer(http.Dir("web/static"))
    http.Handle("GET /static/", http.StripPrefix("/static/", fs))

    // JSON
    http.HandleFunc("GET /search-teams", func(w http.ResponseWriter, r *http.Request) {
        handle.SearchTeams(w, r, pgc)
    })

    // HTML
    http.HandleFunc("GET /teams/", func(w http.ResponseWriter, r *http.Request) {
        handle.Teams(w, r, pgc)
    })
    http.HandleFunc("GET /predict", handle.Predict)

    fmt.Println("Running on 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
