package main

import (
	"routespkg/google_API"

	"github.com/kr/pretty"
)

var cl google_API.ClientData

func main() {
	cl.Origin = "Skopje"
	cl.Destination = "Veles"

	routes := google_API.NewClientData(cl.Origin, cl.Destination)

	pretty.Println(routes.FindRoute())
}
