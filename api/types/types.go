package types

// Mainly just for reading in teams.json right now
type Teams struct {
	Teams []Team `json:"teams"`
}

// Base Team struct, based on Team in api.collegefootballdata.com
type Team struct {
	Id int `json:"id"`
	School string `json:"school"`
	Mascot string `json:"mascot"`
	Abbreviation string `json:"abbreviation"`
	AltName1 string `json:"alt_name1"`
	AltName2 string `json:"alt_name2"`
	AltName3 string `json:"alt_name3"`
	Conference string `json:"conference"`
	Color string `json:"color"`
	AltColor string `json:"alt_color"`
	Logos []string `json:"logos"`
	Twitter string `json:"twitter"`
	Location Location `json:"location"`
}

// Base Location struct, based on Location in api.collegefootballdata.com
type Location struct {
	Id int `json:"venue_id"`
	Name string `json:"name"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
	Country string `json:"country_code"`
	Timezone string `json:"timezone"`
	Lat float64 `json:"latitude"`
	Long float64 `json:"longitude"`
	Elevation string `json:"elevation"`
	Capacity int `json:"capacity"`
	YearConstructed int `json:"year_constructed"`
	Grass bool `json:"grass"`
	Dome bool `json:"dome"`
}

// Small version of Team, just id and name right now
type SMTeam struct {
	Id int `json:"id"`
	School string `json:"school"`
}

// For Teams HTML template
type TeamsTMPLParams struct {
	Team *Team `json:"team"`
	ServerURL string `json:"serverURL"`
}

// Response struct for /search-teams route
type SearchTeamsResponse struct {
	Teams []SMTeam `json:"teams"`
	Error string `json:"error"`
}