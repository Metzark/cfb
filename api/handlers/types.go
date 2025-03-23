package handlers

import pg "github.com/Metzark/cfb/api/pg"

// For Teams HTML template
type TeamsTMPLParams struct {
	Team *pg.Team `json:"team"`
	ServerURL string `json:"serverURL"`
	TeamsList string `json:"teams_list"`
}

// For Query HTML template
type QueryTMPLParams struct {
	ServerURL string `json:"serverURL"`
}

// Response struct for /search-teams route
type SearchTeamsResponse struct {
	Teams []pg.SMTeam `json:"teams"`
	Error string `json:"error"`
}

// Response struct for /custom-query route
type CustomQueryResponse struct {
	Error string `json:"error"`
	Rows []map[string]interface{} `json:"rows"`
 }

 // Request struct for /custom-query route
type CustomQueryRequestBody struct {
	Query string `json:"query"`
}