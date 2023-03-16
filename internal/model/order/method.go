package order

import (
	"fmt"
	"strconv"
	"strings"
)

func (o OrderInquiry) RoutesList() ([]Location, error) {
	locs := make([]Location, 0)

	routes := strings.Split(o.Routes, ";")
	for _, route := range routes {
		loc := strings.Split(route, ",")

		lat, err := strconv.ParseFloat(loc[0], 64)
		if err != nil {
			return nil, err
		}

		lng, err := strconv.ParseFloat(loc[1], 64)
		if err != nil {
			return nil, err
		}

		locs = append(locs, Location{
			Latitude:  lat,
			Longitude: lng,
		})
	}

	return locs, nil
}

func (d Direction) PolylineToStr() string {
	strs := make([]string, 0)
	for _, location := range d.OverviewPolyline {
		strs = append(strs, fmt.Sprintf("%.15f,%.15f", location.Latitude, location.Longitude))
	}

	return strings.Join(strs, ";")
}
