package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	scrapeURL := "https://www.iana.org/domains/reserved"

	collector := colly.NewCollector(colly.AllowedDomains("www.iana.org"))

	collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting => %v\n", r.URL)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while visiting %v\n", err.Error())
	})

	collector.OnHTML("h1", func(h *colly.HTMLElement) {
		fmt.Println("Managed Reserved Domains are =>", h.Text)
	})

	collector.Visit(scrapeURL)
}
