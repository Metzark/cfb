package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	pg "github.com/Metzark/cfb/api/pg"
)

const serverUrl string  = "http://localhost:8080"

// Teams page
func Teams(w http.ResponseWriter, r *http.Request, pgc *pg.PGC){
	var params TeamsTMPLParams

	// Split the url path
	parts := strings.Split(r.URL.Path, "/teams/")

	// Check for valid url path structure
	if len(parts) != 2 {
		http.Error(w, "Path not found", 404)
		return
	}

	// Parse id into integer
	id, err := strconv.Atoi(parts[1])

	// Check that id is a valid integer
	if err != nil {
		http.Error(w, "Team not found", 404)
		return
	}

	// Query team from postgres
	params.Team, err = pg.GetTeamById(pgc, id)

	if err != nil {
		fmt.Println(err)
	}
	
	// Set additional values before filling template
	params.ServerURL = serverUrl

	// Parse the teams html file
	tmpl := template.Must(template.ParseFiles("web/html/teams/index.html"))

	// Set content type for html
	w.Header().Set("Content-Type", "text/html")

	// Write and fill in template
	err = tmpl.Execute(w, params)

	if err != nil {
		fmt.Println(err)
	}
}

// Predict page
func Predict(w http.ResponseWriter, r *http.Request){
	// Parse the predict html file
	tmpl := template.Must(template.ParseFiles("web/html/predict/index.html"))

	// Set content type for html
	w.Header().Set("Content-Type", "text/html")

	// Write and fill in template
	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println(err)
	}
}

// Query page
func Query(w http.ResponseWriter, r *http.Request, pgc *pg.PGC) {
	var params QueryTMPLParams

	// Set additional values before filling template
	params.ServerURL = serverUrl

	// Parse the predict html file
	tmpl := template.Must(template.ParseFiles("web/html/query/index.html"))

	// Set content type for html
	w.Header().Set("Content-Type", "text/html")

	// Write and fill in template
	err := tmpl.Execute(w, params)

<<<<<<< HEAD
	// if err != nil {
	// 	fmt.Println(err)
	// }
=======
	if err != nil {
		fmt.Println(err)
	}
>>>>>>> 5fc5e57 (Update search teams to pull from postgres)
}