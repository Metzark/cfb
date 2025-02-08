package pgc

import (
	"context"
	"fmt"
	"log"
	"os"

	t "github.com/Metzark/cfb/api/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Postgres client
type PGC struct {
	pool *pgxpool.Pool
}

var getTeamByIdQuery string = `
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

// Create a pgc (postgres client) from .env
func CreatePGC() *PGC {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load environment file")
	}

	// Load .env values
	host := os.Getenv("CFB_DB_HOST")
    port := os.Getenv("CFB_DB_PORT")
    user := os.Getenv("CFB_DB_USER")
    password := os.Getenv("CFB_DB_PASSWORD")
    name := os.Getenv("CFB_DB_NAME")
    sslmode := os.Getenv("CFB_DB_SSLMODE")

	// Check for missing .env values
	if host == "" || port == "" || user == "" || password == "" || name == "" || sslmode == "" {
		log.Fatal("Failed loading environment values. Not all of the required values are present.")
    }

	// Create connection string
	connStr := fmt.Sprintf(
        "postgresql://%s:%s@%s:%s/%s?sslmode=%s",
        user, password, host, port, name, sslmode,
    )

	// Create pool
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal("Failed creating connection pool")
	}
	
	// Create postgres client
	pg := PGC{ pool }

	return &pg
}

// Close a pgc (postres client)
func (pgc *PGC) ClosePGC() {
	pgc.pool.Close()
}

// Get a team by their id
func (pgc *PGC) GetTeamById(id int) (*t.Team, error) {
	var pgTeam t.PGTeam
	var team t.Team
	
	// Run query
	teams, err := query[t.PGTeam](pgc, &getTeamByIdQuery, id)

	// Return on error from query
	if err != nil {
		return &team, err
	}

	// Should always be length 1 (if not there should be an err)
	if len(teams) == 1 {
		pgTeam = teams[0]
	}

	// Extract pgtypes to usable values
	t.PGExtract(&pgTeam, &team)

	return &team, nil
}

// Query postgres with the provided client, a parametrized query string, and args
func query[T any](pgc *PGC, q *string, args ...interface{}) ([]T, error) {
	// Array of generic type for the rows returned from the query
	var arr []T

	// Query postgres
	rows, err := pgc.pool.Query(context.Background(), *q, args...)

	// Return query error if there was one
	if err != nil {
		return arr, fmt.Errorf("error executing query: %v", err)
	}

	// Close the rows before return
	defer rows.Close()

	// Collect rows into the generic type array
	arr, err = pgx.CollectRows(rows, pgx.RowToStructByName[T])

	// Return error if there was one when collecting rows
	if err != nil {
		return arr, fmt.Errorf("error collecting rows after query: %v", err)
	}

	return arr, nil
}