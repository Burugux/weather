package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Burugux/weather/geo"
)

type darksky struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Currently struct {
		Time                int     `json:"time"`
		Summary             string  `json:"summary"`
		Icon                string  `json:"icon"`
		PrecipIntensity     float64 `json:"precipIntensity"`
		PrecipProbability   float64 `json:"precipProbability"`
		PrecipType          string  `json:"precipType"`
		Temperature         float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		DewPoint            float64 `json:"dewPoint"`
		Humidity            float64 `json:"humidity"`
		Pressure            float64 `json:"pressure"`
		WindSpeed           float64 `json:"windSpeed"`
		WindGust            float64 `json:"windGust"`
		WindBearing         int     `json:"windBearing"`
		CloudCover          float64 `json:"cloudCover"`
		UvIndex             int     `json:"uvIndex"`
		Visibility          float64 `json:"visibility"`
		Ozone               float64 `json:"ozone"`
	} `json:"currently"`
}

func main() {
	g := geo.Locate()

	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal("API_KEY not set")
	}

	url := fmt.Sprintf("https://api.darksky.net/forecast/%s/%f,%f?units=auto&exclude=hourly,daily", apiKey, g.Latitude, g.Longitude)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("err:", err)
	}
	defer resp.Body.Close()

	var d darksky

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		log.Fatalf("Decoding forecast response failed: %v", err)
	}
	fmt.Println("current location: ", d.Timezone)
	fmt.Println("current temp: ", d.Currently.Temperature)
	fmt.Println("feels like: ", d.Currently.ApparentTemperature)
}
