package google_API

type GoogleRouteResponse struct {
	Legs []*Leg `json:"legs"`
}
type Leg struct {
	Steps         []Steps       `json:"steps"`
	Distance      Distance      `json:"distance"`
	StartLocation StartLocation `json:"start_location"`
	EndLocation   EndLocation   `json:"end_location"`
	StartAddress  string        `json:"start_address"`
	EndAddress    string        `json:"end_address"`
	Duration      Duration      `json:"duration"`
}
type Steps struct {
	Distance      Distance      `json:"distance"`
	StartLocation StartLocation `json:"start_location"`
	EndLocation   EndLocation   `json:"end_location"`
	Steps         interface{}   `json:"steps"`
	Duration      Duration      `json:"duration"`
}
type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}
type StartLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type EndLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Duration struct {
	Value int    `json:"value"`
	Text  string `json:"text"`
}
