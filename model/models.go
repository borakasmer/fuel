package model

import (
	"fmt"
	"strings"
)

type FuelPrice struct {
	Petrol      string
	Diesel      string
	Lpg         string
	City        string
	CurrentDate string
}

type City struct {
	Name string
	Url  string
}

// return a new city instance. also sets the url
func NewCity(cityname string) *City {
	c := &City{Name: cityname}
	c.GenerateURL()
	return c
}

var trCharsToEnChars = map[string]string{
	"ç": "c", "ğ": "g", "ı": "i", "ö": "o", "ş": "s", "ü": "u",
}

func (c *City) GenerateURL() {
	urlTemplate := func(cityname string) string {
		return fmt.Sprintf("https://www.petrolofisi.com.tr/akaryakit-fiyatlari/%s-akaryakit-fiyatlari", cityname)
	}
	loweredCityName := strings.ToLower(c.Name)
	// replace non-english letters with english characters (e.g. ç -> c, ı -> i, ...)
	for tr, en := range trCharsToEnChars {
		loweredCityName = strings.Replace(loweredCityName, tr, en, -1)
	}
	c.Url = urlTemplate(strings.ToLower(loweredCityName))
}
