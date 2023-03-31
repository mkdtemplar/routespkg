package google_API

import (
	"context"

	"googlemaps.github.io/maps"
)

func (cl *ClientData) FindRoute() ([]Route, error) {
	var apiKey = ""

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	r := &maps.DirectionsRequest{
		Origin:      cl.Origin,
		Destination: cl.Destination,
	}

	routes, _, err := c.Directions(context.Background(), r)
	if err != nil {
		return nil, err
	}
	return ToRoutes(routes), nil
}

func ToRoutes(routes []maps.Route) []Route {
	var output []Route
	for _, r := range routes {
		output = append(output, ToRoute(r))
	}
	return output
}

func ToRoute(route maps.Route) Route {
	return Route{
		Summary: route.Summary,
		Legs:    ToLegs(route.Legs),
	}
}

func ToLegs(legs []*maps.Leg) []Leg {
	var output []Leg
	for _, leg := range legs {
		output = append(output, ToLeg(leg))
	}
	return output
}

func ToLeg(leg *maps.Leg) Leg {
	return Leg{
		Steps: ToSteps(leg.Steps),
		Distance: Distance{
			Text:  leg.Distance.HumanReadable,
			Value: leg.Distance.Meters,
		},
		StartLocation: Location{
			Lat: leg.StartLocation.Lat,
			Lng: leg.StartLocation.Lng,
		},
		EndLocation: Location{
			Lat: leg.EndLocation.Lat,
			Lng: leg.EndLocation.Lng,
		},
		StartAddress: leg.StartAddress,
		EndAddress:   leg.EndAddress,
		Duration:     leg.Duration,
	}
}

func ToSteps(steps []*maps.Step) []Steps {
	var output []Steps
	for _, step := range steps {
		output = append(output, ToStep(step))
	}
	return output
}

func ToStep(step *maps.Step) Steps {
	return Steps{
		Distance: Distance{
			Text:  step.Distance.HumanReadable,
			Value: step.Distance.Meters,
		},
		StartLocation: Location{
			Lat: step.StartLocation.Lat,
			Lng: step.StartLocation.Lng,
		},
		EndLocation: Location{
			Lat: step.EndLocation.Lat,
			Lng: step.EndLocation.Lng,
		},
		Steps:    nil,
		Duration: step.Duration,
	}
}
