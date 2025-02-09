package pg

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Postgres client
type PGC struct {
	pool *pgxpool.Pool
}

// PG Team struct, based on Team in api.collegefootballdata.com
type PGTeam struct {
	Id pgtype.Int4 `db:"team_id"`
	School pgtype.Text `db:"team_school"`
	Mascot pgtype.Text `db:"team_mascot"`
	Abbreviation pgtype.Text `db:"team_abbreviation"`
	AltName1 pgtype.Text `db:"team_alt_name1"`
	AltName2 pgtype.Text `db:"team_alt_name2"`
	AltName3 pgtype.Text `db:"team_alt_name3"`
	Classification pgtype.Text `db:"team_classification"`
	Color pgtype.Text `db:"team_color"`
	AltColor pgtype.Text `db:"team_alt_color"`
	Logo pgtype.Text `db:"team_logo"`
	Twitter pgtype.Text `db:"team_twitter"`
	VenueId pgtype.Int4 `db:"venue_id"`
	VenueName pgtype.Text `db:"venue_name"`
	VenueCity pgtype.Text `db:"venue_city"`
	VenueState pgtype.Text `db:"venue_state"`
	VenueZip pgtype.Text `db:"venue_zip"`
	VenueCountry pgtype.Text `db:"venue_country_code"`
	VenueTimezone pgtype.Text `db:"venue_timezone"`
	VenueLocationX pgtype.Float8 `db:"venue_location_x"`
	VenueLocationY pgtype.Float8 `db:"venue_location_y"`
	VenueElevation pgtype.Float8 `db:"venue_elevation"`
	VenueCapacity pgtype.Int4 `db:"venue_capacity"`
	VenueYearConstructed pgtype.Int4 `db:"venue_year_constructed"`
	VenueGrass pgtype.Bool `db:"venue_grass"`
	VenueDome pgtype.Bool `db:"venue_dome"`
	ConferenceId pgtype.Int4 `db:"conference_id"`
	ConferenceName pgtype.Text `db:"conference_name"`
}

// PG Team struct, based on Team in api.collegefootballdata.com
type Team struct {
	Id int32 `json:"team_id"`
	School string `json:"team_school"`
	Mascot string `json:"team_mascot"`
	Abbreviation string `json:"team_abbreviation"`
	AltName1 string `json:"team_alt_name1"`
	AltName2 string `json:"team_alt_name2"`
	AltName3 string `json:"team_alt_name3"`
	Classification string `json:"team_classification"`
	Color string `json:"team_color"`
	AltColor string `json:"team_alt_color"`
	Logo string `json:"team_logo"`
	Twitter string `json:"team_twitter"`
	VenueId int32 `json:"venue_id"`
	VenueName string `json:"venue_name"`
	VenueCity string `json:"venue_city"`
	VenueState string `json:"venue_state"`
	VenueZip string `json:"venue_zip"`
	VenueCountry string `json:"venue_country_code"`
	VenueTimezone string `json:"venue_timezone"`
	VenueLocationX float64 `json:"venue_location_x"`
	VenueLocationY float64 `json:"venue_location_y"`
	VenueElevation float64 `json:"venue_elevation"`
	VenueCapacity int32 `json:"venue_capacity"`
	VenueYearConstructed int32 `json:"venue_year_constructed"`
	VenueGrass bool `json:"venue_grass"`
	VenueDome bool `json:"venue_dome"`
	ConferenceId int32 `json:"conference_id"`
	ConferenceName string `json:"conference_name"`
}

type PGSMTeam struct {
	Id int `db:"id"`
	School string `db:"school"`
}

type SMTeam struct {
	Id int `json:"id"`
	School string `json:"school"`
}

// PG Venue struct, based on Location in api.collegefootballdata.com
type PGVenue struct {
	Id pgtype.Int4 `db:"venue_id"`
	Name pgtype.Text `db:"venue_name"`
	City pgtype.Text `db:"venue_city"`
	State pgtype.Text `db:"venue_state"`
	Zip pgtype.Text `db:"venue_zip"`
	Country pgtype.Text `db:"venue_country_code"`
	Timezone pgtype.Text `db:"venue_timezone"`
	LocationX pgtype.Float8 `db:"venue_location_x"`
	LocationY pgtype.Float8 `db:"venue_location_y"`
	Elevation pgtype.Float8 `db:"venue_elevation"`
	Capacity pgtype.Int4 `db:"venue_capacity"`
	YearConstructed pgtype.Int4 `db:"venue_year_constructed"`
	Grass pgtype.Bool `db:"venue_grass"`
	Dome pgtype.Bool `db:"venue_dome"`
}

// PG Conference struct based on Conference in api.collegefootballdata.com
type PGConference struct {
	Id pgtype.Int4 `db:"conference_id"`
	Name pgtype.Text `db:"conference_name"`
}

// Extract actual values (into s2) from a struct with pgtype values (from s1)
func PGExtract[T1 any, T2 any](s1 *T1, s2 *T2) {

	// Values from struct 1
	s1Value := reflect.ValueOf(s1).Elem()

	// Types from struct 1
	s1Type := s1Value.Type()

	// Reflect struct 2 so that the values can be set
	s2Value := reflect.ValueOf(s2).Elem()

	// For each field in struct 1
	for i := 0; i < s1Value.NumField(); i++ {
		// Get field from struct 1
		field := s1Type.Field(i)

		// Get value from struct 1
		value := s1Value.Field(i)

		// Get the corresponding field for struct 2
		s2Field := s2Value.FieldByName(field.Name)

		// Go through pgtypes, check for valid values and set field for struct 2
		switch v := value.Interface().(type) {
		case pgtype.Text:
			if v.Status == pgtype.Present {
				s2Field.Set(reflect.ValueOf(v.String))
			} else {
				s2Field.Set(reflect.Zero(s2Field.Type()))
			}
		case pgtype.Int4:
			if v.Status == pgtype.Present {
				s2Field.Set(reflect.ValueOf(v.Int))
			} else {
				s2Field.Set(reflect.Zero(s2Field.Type()))
			}
		case pgtype.Float8:
			if v.Status == pgtype.Present {
				s2Field.Set(reflect.ValueOf(v.Float))
			} else {
				s2Field.Set(reflect.Zero(s2Field.Type()))
			}
		case pgtype.Bool:
			s2Field.Set(reflect.ValueOf(v.Bool))
		default:
			if s2Field.Type() == value.Type() {
				s2Field.Set(value)
			}
		}
	}
}


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

// Query postgres with the provided client, a parametrized query string, and args
func query[T any](pgc *PGC, q string, args ...interface{}) ([]T, error) {
	// Array of generic type for the rows returned from the query
	var arr []T

	// Query postgres
	rows, err := pgc.pool.Query(context.Background(), q, args...)

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