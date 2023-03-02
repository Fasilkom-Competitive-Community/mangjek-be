package _map

import (
	"context"
	"errors"
	"fmt"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	"googlemaps.github.io/maps"
	"time"
)

type MapCalculator struct {
	client *maps.Client
}

func NewMapCalculator(apiKey string) (*MapCalculator, error) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &MapCalculator{client: client}, nil
}

func (m *MapCalculator) CalculateDirection(ctx context.Context, origin oModel.Location, destination oModel.Location) (oModel.Direction, error) {
	// Format float with 15 digits after comma
	org := fmt.Sprintf("%.15f|%.15f", origin.Latitude, origin.Longitude)
	dst := fmt.Sprintf("%.15f|%.15f", destination.Latitude, destination.Longitude)

	request := &maps.DirectionsRequest{
		Origin:        org,
		Destination:   dst,
		Mode:          maps.TravelModeDriving,
		DepartureTime: fmt.Sprintf("%d", time.Now().Unix()),
		Language:      "ID",
		Region:        "ID",
	}

	routes, _, err := m.client.Directions(ctx, request)
	if err != nil {
		return oModel.Direction{}, err
	}

	if len(routes) == 0 {
		return oModel.Direction{}, errors.New("MAP_COMMON.ROUTE_IS_NOT_FOUND")
	}

	decode, err := routes[0].OverviewPolyline.Decode()
	if err != nil {
		return oModel.Direction{}, err
	}

	dr := oModel.Direction{
		Distance: int32(routes[0].Legs[0].Distance.Meters),
		Duration: int32(routes[0].Legs[0].Duration.Seconds()),
		Origin: oModel.Location{
			Address:   routes[0].Legs[0].StartAddress,
			Latitude:  routes[0].Legs[0].StartLocation.Lat,
			Longitude: routes[0].Legs[0].StartLocation.Lng,
		},
		Destination: oModel.Location{
			Address:   routes[0].Legs[0].EndAddress,
			Latitude:  routes[0].Legs[0].EndLocation.Lat,
			Longitude: routes[0].Legs[0].EndLocation.Lng,
		},
		OverviewPolyline: make([]oModel.Location, 0),
	}

	for _, polyline := range decode {
		dr.OverviewPolyline = append(dr.OverviewPolyline, oModel.Location{
			Latitude:  polyline.Lat,
			Longitude: polyline.Lng,
		})
	}

	return dr, nil
}
