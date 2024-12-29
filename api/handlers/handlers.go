package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	io "github.com/Metzark/cfb/api/io"
	t "github.com/Metzark/cfb/api/types"
)

var serverUrl string  = "http://localhost:8080"

var teams []t.Team = io.GetTeams()


// Teams page HTML
func Teams(w http.ResponseWriter, r *http.Request){
	var params t.TeamsTMPLParams

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

	// Match team ids
	for i := range len(teams) {
		if id == teams[i].Id {
			params.Team = &teams[i]
		}
	}

	// Set additional values before filling template
	params.ServerURL = serverUrl

	// Parse the teams html file
	tmpl := template.Must(template.ParseFiles("web/html/teams/index.html"))

	// Write and fill in template
	err = tmpl.Execute(w, params)

	if err != nil {
		fmt.Println(err)
	}
}

// Search teams JSON
func SearchTeams(w http.ResponseWriter, r *http.Request){
	var res t.SearchTeamsResponse

	// Set response content type for json
	w.Header().Set("Content-Type", "application/json")

	// Get request query
	query := r.URL.Query()

	// Should only be 1 value for search
	if len(query["search"]) != 1 {
		res.Error = "The search query must have 1 value..."
		json.NewEncoder(w).Encode(res)
		return
	}

	search := query["search"][0]

	var match bool

	// Substring match to get searched teams
	for i := range len(teams) {
		// Match the search to either the school or one of the alt names
		match = (strings.HasPrefix(strings.ToLower(teams[i].School), strings.ToLower(search)) || 
			strings.HasPrefix(strings.ToLower(teams[i].AltName1), strings.ToLower(search)) ||
			strings.HasPrefix(strings.ToLower(teams[i].AltName2), strings.ToLower(search)) ||
			strings.HasPrefix(strings.ToLower(teams[i].AltName2), strings.ToLower(search)))

		// If search match, add to response
		if match {
			res.Teams = append(res.Teams, t.SMTeam{ Id: teams[i].Id, School: teams[i].School })
		}
	}

	// Write json 
	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		fmt.Println(err)
	}
}