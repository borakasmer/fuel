package parser

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/borakasmer/fuel/core"
)

/*
doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))

    if err != nil {
        log.Fatal(err)
    }

    var words []string

    sel1 := doc.Find("li:first-child, li:last-child")
    sel2 := doc.Find("li:nth-child(3), li:nth-child(7)")

    sel1.Union(sel2).Each(func(_ int, sel *goquery.Selection) {
        words = append(words, sel.Text())
    })

    fmt.Println(words)
}
The example combines two selections.

sel1 := doc.Find("li:first-child, li:last-child")
The first selection contains the first and the last element.

sel2 := doc.Find("li:nth-child(3), li:nth-child(7)")
*/
func ParseWeb(url string) (core.String, core.String, core.String) {
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
			data := doc.Find("#fuelPricesTableDesktop tbody tr:nth-child(1)")
			data.Each(func(i int, s *goquery.Selection) {
				petrol = s.Find(".data-cell:nth-child(2)").Text()
				diesel = s.Find(".data-cell:nth-child(3)").Text()
				lpg = s.Find(".data-cell:nth-child(5)").Text()
			})
		}
	}
	return core.String{Value: petrol}, core.String{Value: diesel}, core.String{Value: lpg}
}
