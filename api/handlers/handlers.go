package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	pg "github.com/Metzark/cfb/api/pg"
)

var serverUrl string  = "http://localhost:8080"

// For Teams HTML template
type TeamsTMPLParams struct {
	Team *pg.Team `json:"team"`
	ServerURL string `json:"serverURL"`
}


// Response struct for /search-teams route
type SearchTeamsResponse struct {
	Teams []pg.SMTeam `json:"teams"`
	Error string `json:"error"`
}

// Teams page HTML
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

// Predict page HTML
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

// Search teams JSON
func SearchTeams(w http.ResponseWriter, r *http.Request, pgc *pg.PGC){
	var res SearchTeamsResponse
	var err error

	// // Set response content type for json
	// w.Header().Set("Content-Type", "application/json")

	// // Get request query
	// query := r.URL.Query()

	// // Should only be 1 value for search
	if len(query["search"]) != 1 {
		res.Error = "The search query must have 1 value..."
		json.NewEncoder(w).Encode(res)
		return
	}

	// search := query["search"][0]

	
	res.Teams, err = pg.SearchTeamsByName(pgc, search)

	if err != nil {
		fmt.Println(err)
		res.Error = err.Error()
	}

	// // Write json using writer (this sends a response)
	err = json.NewEncoder(w).Encode(res)

	// if err != nil {
	// 	fmt.Println(err)
	// }
}