package google_API

import "time"

type GoogleRouteResponse struct {
	Legs []Leg `json:"legs"`
}

type Route struct {
	Summary string
	Legs    []Leg
}

type Leg struct {
	Steps         []Steps       `json:"steps"`
	Distance      Distance      `json:"distance"`
	StartLocation Location      `json:"start_location"`
	EndLocation   Location      `json:"end_location"`
	StartAddress  string        `json:"start_address"`
	EndAddress    string        `json:"end_address"`
	Duration      time.Duration `json:"duration"`
}
type Steps struct {
	Distance      Distance      `json:"distance"`
	StartLocation Location      `json:"start_location"`
	EndLocation   Location      `json:"end_location"`
	Steps         interface{}   `json:"steps"`
	Duration      time.Duration `json:"duration"`
}
type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Duration struct {
	Value int    `json:"value"`
	Text  string `json:"text"`
}
