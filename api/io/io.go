package io

import (
	"encoding/json"
	"io"
	"log"
	"os"

	t "github.com/Metzark/cfb/api/types"
)

// Get the teams from data/teams.json
// WILL FATAL IF data/teams.json CANT BE READ IN!
func GetTeams() []t.Team {
	// Open teams.json file
	file, err := os.Open("data/teams.json")

	if err != nil {
		log.Fatal("Failed to open teams.json file...\n", err)
	}
	
	// Read from teams.json file
	teamsBytes, err := io.ReadAll(file)

	if err != nil {
		log.Fatal("Failed to read teams.json file...\n", err)
	}

	// Close the teams.json file
	file.Close()

	var teams t.Teams

	// Move data into the teams array
	err = json.Unmarshal(teamsBytes, &teams)

	if err != nil {
		log.Fatal("Failed to unmarshall teams.json file...\n", err)
	}

	return teams.Teams
}