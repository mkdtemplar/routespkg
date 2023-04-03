package google_API

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GasStations() (string, error) {
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=41.99646,%2021.43141&radius=50000&type=gas_station&key=INSERT_API_KEY"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "nil", err
	}

	mapResponse := GasStationsObject{}

	err = json.Unmarshal(body, &mapResponse)
	for _, i := range mapResponse.Results {
		fmt.Println(i.BusinessStatus)
		fmt.Println(i.Name)
		fmt.Println(i.Geometry.Location)
		fmt.Println(i.OpeningHours.OpenNow)
	}

	fmt.Println(mapResponse)

	return string(body), nil
}
