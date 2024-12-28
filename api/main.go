package main

import (
	"fmt"
	"log"
	"net/http"

	handle "github.com/Metzark/cfb/api/handlers"
)


func main(){
    http.HandleFunc("GET /", handle.Teams)

    fmt.Println("Running on 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

