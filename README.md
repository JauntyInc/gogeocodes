# GeoCodes

The alpha version of the [https://geo.codes](https://geo.codes) API service. See API documentation here. 10,000 free queries per month.

## Installation & Usage

### Requirements

To install, simply use `go get`

```go
go get github.com/JauntyInc/gogeocodes
```

Check out the `example/` directory to see a working program.

## Getting Started

Please follow the installation procedure and then run the following:

Note that this is a reproduction of the full working program, which can be found in the example/
directory

```go
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
```

## API Endpoints

All URIs are relative to *https://api.geo.codes*

Client | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*client.Address* | **V1AddressGeocode** | **GET** /v1/address/geocode | Geocode an unstructured address string
*client.Address* | **V1AddressStructuredGeocode** | **GET** /v1/address/structured_geocode | Geocode structured addresses
*client.Auth* | **V1Test** | **GET** /v1/test | Ping function that tests the API Key
*client.Operations* | **Ping** | **GET** /ping | Ping the service without credentials
*client.ZIPCode* | **V1AddressReverseZipCode** | **GET** /v1/address/reverse_zip_code | Convert a coordinate to a ZIP Code
*client.ZIPCode* | **V1AddressZipCode** | **GET** /v1/address/zip_code | Convert a ZIP Code to a coordinate

## Getting help

If you need help with this library, email [help@geo.codes](mailto:help@geo.codes) and we will
be happy to assist you.
