package pg

import (
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
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
	TeamsList pgtype.JSONB  `db:"teams_list"`
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
	TeamsList []SMTeam `json:"teams_list"`
}

type PGSMTeam struct {
	Id int `db:"team_id"`
	School string `db:"team_school"`
	Logo string `db:"team_logo"`
}

type SMTeam struct {
	Id int `json:"team_id"`
	School string `json:"team_school"`
	Logo string `json:"team_logo"`
	Mascot string `json:"team_mascot"`
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