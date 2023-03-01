package main

import (
	"context"
	"fmt"
	"googlemaps.github.io/maps"
	"strconv"
	"time"
)

func main() {
	client, err := maps.NewClient(maps.WithAPIKey("AIzaSyDIcTGa61FUTuSvoN1W5oRaLlF3K-Bfbmo"))
	if err != nil {
		panic(err)
	}

	request := &maps.DirectionsRequest{
		Origin:        "-0.955656130989188,114.33822509249502",
		Destination:   "39.17342299237161,-95.6368425602725",
		Mode:          maps.TravelModeDriving,
		TrafficModel:  maps.TrafficModelBestGuess,
		DepartureTime: strconv.Itoa(int(time.Now().Unix())),
		Language:      "ID",
		Region:        "ID",
	}

	directions, _, err := client.Directions(context.Background(), request)
	if err != nil {
		panic(err)
	}

	for _, direction := range directions {
		//fmt.Printf("di, %+v\n", direction)
		fmt.Println("AAAAAAAAAA")
		decode, err := direction.OverviewPolyline.Decode()
		if err != nil {
			panic(err)
		}
		fmt.Println(decode)
		for _, leg := range direction.Legs {
			fmt.Printf("leg, %+v\n", leg)
		}
	}

	//for _, waypoint := range i {
	//	fmt.Printf("waypoint, %+v\n", waypoint)
	//}
}
