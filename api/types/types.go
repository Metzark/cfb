package types

type Teams struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	Id int `json:"id"`
	School string `json:"school"`
	Mascot string `json:"mascot"`
	Abbreviation string `json:"abbreviation"`
	Conference string `json:"conference"`
	Color string `json:"color"`
	AltColor string `json:"alt_color"`
	Logos []string `json:"logos"`
	Twitter string `json:"twitter"`
	Location Location `json:"location"`
}

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