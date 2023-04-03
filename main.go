package main

import (
	"fmt"
	"log"
	"routespkg/google_API"

	"github.com/kr/pretty"
)

var cl google_API.ClientData

func main() {

	cl.Origin = "Skopje"
	cl.Destination = "Veles"

	routes := google_API.NewClientData(cl.Origin, cl.Destination)

	gasStations, err := google_API.GasStations()
	if err != nil {
		log.Println(err)
	}
	/*
		fmt.Println("Distance matrix: ")
		google_API.DistanceMatrix()
	*/

	fmt.Println(gasStations)

	pretty.Println(routes.FindRoute())

	//google_API.DistanceMatrix()
}
