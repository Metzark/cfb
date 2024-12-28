package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	io "github.com/Metzark/cfb/api/io"
	t "github.com/Metzark/cfb/api/types"
)

var teams []t.Team = io.GetTeams()


func Teams(w http.ResponseWriter, r *http.Request){
	var team *t.Team

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
			team = &teams[i]
		}
	}

	tmpl := template.Must(template.ParseFiles("web/html/teams/index.html"))
	tmpl.Execute(w, team)
}
