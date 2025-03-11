package pg

import "fmt"


const getTeamByIdQuery string = `
	SELECT 
		t.id AS team_id, t.school AS team_school, t.mascot AS team_mascot, t.abbreviation AS team_abbreviation, t.alt_name1 AS team_alt_name1,
		t.alt_name2 AS team_alt_name2, t.alt_name3 AS team_alt_name3, t.classification AS team_classification,
		t.color AS team_color, t.alt_color AS team_alt_color, t.twitter AS team_twitter, t.logo AS team_logo,
		v.id AS venue_id, v.name AS venue_name, v.city AS venue_city, v.state as venue_state, v.zip AS venue_zip,
		v.country_code AS venue_country_code, v.location_x AS venue_location_x, v.location_y AS venue_location_y,
		v.elevation AS venue_elevation, v.year_constructed AS venue_year_constructed, v.capacity AS venue_capacity,
		v.dome AS venue_dome, v.grass AS venue_grass, v.timezone AS venue_timezone,
		c.id AS conference_id, c.name AS conference_name
	FROM cfb.teams t
	LEFT JOIN cfb.conferences c ON c.id = t.conference_id
	LEFT JOIN cfb.venues v ON v.id = t.home_venue_id
	WHERE t.id = $1;
`

// Get a team by their id
func GetTeamById(pgc *PGC, id int) (*Team, error) {
	var team Team
	
	// Run parametrized query
	pgTeams, err := query[PGTeam](pgc, getTeamByIdQuery, id)

	// Return on error from query
	if err != nil {
		return &team, err
	}

	// Should always be length 1 (if not there should be an err)
	if len(pgTeams) == 1 {
		// Extract pgtypes to usable values
		PGExtract(&pgTeams[0], &team)
	}

	return &team, nil
}

const searchTeamsByNameQuery string = `
	SELECT id, school
	FROM cfb.teams t
	WHERE
		t.classification ILIKE 'fbs'
		AND (
			t.school ILIKE $1
			OR t.alt_name1 ILIKE $1
			OR t.alt_name2 ILIKE $1
			OR t.alt_name3 ILIKE $1);
`

// Search teams by name (school name and alternative names)
func SearchTeamsByName(pgc *PGC, name string) ([]SMTeam, error) {

	search := fmt.Sprintf("%%%s%%", name)

	// Run parametrized query
	pgSMTeams, err := query[PGSMTeam](pgc, searchTeamsByNameQuery, search)

	smTeams := make([]SMTeam, len(pgSMTeams))

	// Return on error from search
	if err != nil {
		return smTeams, err
	}

	// Extract pgtypes to usable values
	for i := 0; i < len(pgSMTeams); i++ {
		PGExtract(&pgSMTeams[i], &smTeams[i])
	}

	return smTeams, nil
}

func ExecuteCustomQuery(pgc *PGC, query string) ([]map[string]interface{}, error) {
	var rows []map[string]interface{}

	rows, err := customQuery(pgc, query)

	return rows, err
}