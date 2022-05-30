package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"io/ioutil"
)

func Colly2() {
	col := colly.NewCollector()

	col.OnHTML("div", func(element *colly.HTMLElement) {
		fmt.Println(element)
	})

	col.Visit("http://go-colly.org/")
}

func ScrapeHref() *colly.Collector {
	col := colly.NewCollector()
	col.OnHTML("a[href]", func(el *colly.HTMLElement) {
		el.Request.Visit(el.Attr("href"))
	})
	col.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	return col
}

func CollyFromFile(fl string) {
	file, err := ioutil.ReadFile(fl)
	if err != nil {
		return
	}
	reader := bytes.NewReader(file)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
	}
	//divs := doc.Find("div")
	tb := doc.Find("#tab1")
	fmt.Println(tb)
	var res string
	var ndd html.Node
	err = colly.UnmarshalHTML(ndd, tb)
	if err != nil {
		return
	}
	fmt.Println(res)
}
