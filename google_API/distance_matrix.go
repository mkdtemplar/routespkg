package google_API

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func DistanceMatrix() {
	distanceMatrix := DistanceMatrixObject{}

	url := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=41.99646,21.43141&destinations=41.9827805,21.4521448&key=INSERT_API_KEY"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &distanceMatrix)
	if err != nil {
		log.Println("Cannot unmarshal response")
	}

	fmt.Println(distanceMatrix)
}
