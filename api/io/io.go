package io

import (
	"encoding/json"
	"io"
	"os"
)

type Team struct {
	Id int32 `json:"id"`
	School string `json:"school"`
	Mascot string `json:"mascot"`
	Abbreviation string `json:"abbreviation"`
	Conference string `json:"conference"`
	Color string `json:"color"`
	AltColor string `json:"alt_color"`
	Logos []string `json:"logos"`
	Twitter string `json:"twitter"`
	Location location `json:"location"`
}

type location struct {
	Id int32 `json:"venue_id"`
	Name string `json:"name"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
	Country string `json:"country_code"`
	Timezone string `json:"timezone"`
	Lat float64 `json:"latitude"`
	Long float64 `json:"longitude"`
	Elevation float64 `json:"elevation"`
	Capacity int32 `json:"capactiy"`
	YearConstructed int32 `json:"year_constructed"`
	Grass bool `json:"grass"`
	Dome bool `json:"dome"`
}

func getTeams() (*[]Team, error) {
	// Open teams.json file
	file, err := os.Open("data/teams.json")

	if err != nil {
		return nil, err
	}
	
	// Read from teams.json file
	teamsBytes, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	// Close the teams.json file
	file.Close()

	var teams []Team

	// Move data into the teams array
	json.Unmarshal(teamsBytes, &teams)

	return &teams, nil
}