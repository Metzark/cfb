package types

import (
	"reflect"

	"github.com/jackc/pgx/v5/pgtype"
)

// Mainly just for reading in teams.json right now
type Teams struct {
	Teams []Team `json:"teams"`
}

// Base Team struct, based on Team in api.collegefootballdata.com
// type Team struct {
// 	Id int `json:"id"`
// 	School string `json:"school"`
// 	Mascot string `json:"mascot"`
// 	Abbreviation string `json:"abbreviation"`
// 	AltName1 string `json:"alt_name1"`
// 	AltName2 string `json:"alt_name2"`
// 	AltName3 string `json:"alt_name3"`
// 	Conference string `json:"conference"`
// 	Color string `json:"color"`
// 	AltColor string `json:"alt_color"`
// 	Logos []string `json:"logos"`
// 	Twitter string `json:"twitter"`
// 	Location Location `json:"location"`
// }

// Base Location struct, based on Location in api.collegefootballdata.com
// type Location struct {
// 	Id int `json:"venue_id"`
// 	Name string `json:"name"`
// 	City string `json:"city"`
// 	State string `json:"state"`
// 	Zip string `json:"zip"`
// 	Country string `json:"country_code"`
// 	Timezone string `json:"timezone"`
// 	Lat float64 `json:"latitude"`
// 	Long float64 `json:"longitude"`
// 	Elevation string `json:"elevation"`
// 	Capacity int `json:"capacity"`
// 	YearConstructed int `json:"year_constructed"`
// 	Grass bool `json:"grass"`
// 	Dome bool `json:"dome"`
// }

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
			if v.Valid {
				s2Field.Set(reflect.ValueOf(v.String))
			} else {
				s2Field.Set(reflect.Zero(s2Field.Type()))
			}
		case pgtype.Int4:
			if v.Valid {
				s2Field.Set(reflect.ValueOf(v.Int32))
			} else {
				s2Field.Set(reflect.Zero(s2Field.Type()))
			}
		case pgtype.Float8:
			if v.Valid {
				s2Field.Set(reflect.ValueOf(v.Float64))
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