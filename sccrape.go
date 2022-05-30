package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Colly2() {
	col := colly.NewCollector()
	col.OnHTML("div", func(element *colly.HTMLElement) {
		fmt.Println(element)
	})

	col.Visit("http://go-colly.org/")
}

func scrapeHref() *colly.Collector {
	col := colly.NewCollector()
	col.OnHTML("a[href]", func(el *colly.HTMLElement) {
		el.Request.Visit(el.Attr("href"))
	})
	col.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	return col
}
