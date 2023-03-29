package google_API

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

func (cl *ClientData) FindRoute() []maps.Route {
	var apiKey = ""

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil
	}

	r := &maps.DirectionsRequest{
		Origin:      cl.Origin,
		Destination: cl.Destination,
	}

	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)

	}

	for _, i := range route {
		for _, j := range i.Legs {
			fmt.Println("Total distance: ", j.Distance)
			for l, k := range j.Steps {
				fmt.Println("Step: ", l)
				fmt.Println(k.StartLocation)
				fmt.Println(k.EndLocation)
				fmt.Println("Distance: ", k.Distance.HumanReadable)
				fmt.Println("Time duration of the step: ", k.Duration)
				fmt.Println("--------------------")
			}
		}
	}

	return route
}
