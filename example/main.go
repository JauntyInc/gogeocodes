package main

import (
	"fmt"
	"time"

	gogeocodes "github.com/JauntyInc/gogeocodes/client"
	"github.com/JauntyInc/gogeocodes/client/address"
	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	// Create and configure the client
	client := gogeocodes.Default
	auth := httptransport.APIKeyAuth("Authorization", "header", "YOUR_KEY_HERE")
	params := &address.V1AddressGeocodeParams{
		Q: "111 8th Ave, New York, NY 10011",
	}
	params.SetTimeout(time.Second * 2)

	// Send the query
	result, err := client.Address.V1AddressGeocode(params, auth)

	// Handle any errors. If it's a 401, create a key at https://geo.codes/account/api
	if err != nil {
		fmt.Printf("Error fetching geocoded address data: [%s]\n", err.Error())
		return
	}

	// Print the results!
	coordinate := result.Payload.Coordinate
	fmt.Printf(
		"Latitude: [%f] Longitude: [%f]\n", *coordinate.Latitude, *coordinate.Longitude,
	)
}
