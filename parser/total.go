package parser

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/borakasmer/fuel/core"
)

const TotalAnkaraURL = "http://195.216.232.22/product_price.asp?cityID=7"
const TotalIzmirURL = "http://195.216.232.22/product_price.asp?cityID=33"
const TotalIstanbulURL = "http://195.216.232.22/product_price.asp?cityID=32"

func ParseWebTotal(url string) (core.String, core.String, core.String) {
	var petrol = ""
	var diesel = ""
	var lpg = ""

	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			data := doc.Find("table[border='1'] tbody tr")
			data.EachWithBreak(func(i int, s *goquery.Selection) bool {
				var merkez = strings.TrimSpace(s.Find("td:nth-child(1)").Text())
				merkez = strings.Trim(merkez, "-ANADOLU")

				if merkez == "MERKEZ" {
					petrol = strings.Replace(s.Find("td:nth-child(2)").Text(), ",", ".", -1)
					diesel = strings.Replace(s.Find("td:nth-child(4)").Text(), ",", ".", -1)
					lpg = strings.Replace(s.Find("td:nth-child(9)").Text(), ",", ".", -1)
					return false // isimiz bitti. cikalim!
				}

				return true
			})
		}
	}

	return core.String{petrol}, core.String{diesel}, core.String{lpg}
}
