package parser

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/borakasmer/fuel/Model"
)

const _DATE_FORMAT_STRING = "2006.01.02 15:04:05"

type response struct {
	K95   string
	Mot50 string
	PoGaz string
}

func ParseWeb(cityName string) Model.FuelPrice {
	// JSON Keys for Renaming
	// K95 -> Benzin
	// Mot50 -> Mazot
	// PoGaz -> LPG

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Get("https://www.petrolofisi.com.tr/posvc/fiyat/guncel?il=" + cityName + "&ilce=" + cityName)
	if err != nil {
		log.Fatal(err)
	}

	var fuelPricesSlice []response
	defer res.Body.Close()

	if res.StatusCode == 200 {
		decoder := json.NewDecoder(res.Body)
		err := decoder.Decode(&fuelPricesSlice)
		if err != nil {
			log.Fatal("Error on parsing API response on City: ", cityName, "\n", err)
		}
	}

	fuelPrices := fuelPricesSlice[0]

	return Model.FuelPrice{
		City:        cityName,
		Diesel:      fuelPrices.Mot50,
		Petrol:      fuelPrices.K95,
		Lpg:         fuelPrices.PoGaz,
		CurrentDate: time.Now().Format(_DATE_FORMAT_STRING),
	}
}
