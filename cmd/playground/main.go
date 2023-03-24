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
	directionMap()
	//paymentQRISXendit()
	//calculatePrice()
}

func calculatePrice() {
	var price int64
	//distance := 2200
	//distance := 5100
	distance := 5400
	if distance <= 3_000 {
		price = 5_000
	} else {
		price = int64(distance) * 2
		remainder := price % 500
		if remainder != 0 {
			price = price - remainder + 500
		}
	}

	fmt.Println(price)
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
		Origin:        "-3.210136297674036,104.65119733043355",
		Destination:   "-3.2202908430389106,104.65190274763013",
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
