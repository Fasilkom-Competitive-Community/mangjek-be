package main

import (
	"context"
	"fmt"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/qrcode"
	"googlemaps.github.io/maps"
	"log"
	"strconv"
	"time"
)

func main() {
	//directionMap()
	paymentQRISXendit()
}

func paymentQRISXendit() {
	xendit.Opt.SecretKey = "xnd_development_ye78i7MafcV7CF1UPhdFS6m4NjfzXfqtYjlQ9KDiJ3fNMZXG3AFdYfStgtbpZ"
	//createData := qrcode.CreateQRCodeParams{
	//	ExternalID:  "hahayy",
	//	Type:        xendit.DynamicQRCode,
	//	CallbackURL: "https://httpdump.app/dumps/15ef4794-3057-47e7-bb78-51778a5058a4",
	//	Amount:      15_000,
	//}
	//
	//resp, err := qrcode.CreateQRCode(&createData)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("created QR code: %+v\n", resp)

	resp, err := qrcode.GetQRCode(&qrcode.GetQRCodeParams{
		ExternalID: "hahayy",
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("retrieved QR code: %+v\n", resp)
}

func directionMap() {
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
