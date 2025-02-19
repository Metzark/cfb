package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	pg "github.com/Metzark/cfb/api/pg"
)

// Search teams
func SearchTeams(w http.ResponseWriter, r *http.Request, pgc *pg.PGC){
	var res SearchTeamsResponse
	var err error

	// Set response content type for json
	w.Header().Set("Content-Type", "application/json")

	// Get request query
	query := r.URL.Query()

	// // Should only be 1 value for search
	if len(query["search"]) != 1 {
		res.Error = "The search query must have 1 value..."
		json.NewEncoder(w).Encode(res)
		return
	}

	search := query["search"][0]

	res.Teams, err = pg.SearchTeamsByName(pgc, search)

	if err != nil {
		fmt.Println(err)
		res.Error = err.Error()
	}

	// Write json using writer (this sends a response)
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		fmt.Println(err)
	}
}

// Run customer query (for /query)
func CustomQuery(w http.ResponseWriter, r *http.Request, pgc *pg.PGC) {
	var res CustomQueryResponse
	var body CustomQueryRequestBody

	// Set response content type for json
	w.Header().Set("Content-Type", "application/json")

	// Get query from request body
	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		fmt.Println(err)
		res.Error = err.Error()
	}

	// Execute the custom query
	res.Rows, err = pg.ExecuteCustomQuery(pgc, body.Query)

	if err != nil {
		fmt.Println(err)
		res.Error = err.Error()
	}

	// Write json using writer
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		fmt.Println(err)
	}
}