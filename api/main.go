package main

import (
	"fmt"
	"log"
	"net/http"

	handle "github.com/Metzark/cfb/api/handlers"
)


func main(){

    fs := http.FileServer(http.Dir("web/static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("GET /teams/", handle.Teams)

    fmt.Println("Running on 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

