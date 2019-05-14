package geo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//Locate returns the lat and lng of the user
func Locate() GeoLocation {
	resp, err := http.Get("https://ifconfig.co/json")

	if err != nil {
		fmt.Println("err: ", err)
	}
	defer resp.Body.Close()

	var g GeoLocation

	if err := json.NewDecoder(resp.Body).Decode(&g); err != nil {
		log.Fatal(err)
	}
	return g
}
