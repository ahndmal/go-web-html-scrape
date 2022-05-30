package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"io/ioutil"
	"net/http"
)

func main() {
	//CollyFromFile("index.html")
	file, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	}
	reader := bytes.NewReader(file)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
	}
	//divs := doc.Find("div")
	tb := doc.Find("#tab1")
	fmt.Println(tb.Text())

	res := make([]string, 0)
	//var ndd html.Node
	err = colly.UnmarshalHTML(res, tb)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func FromUrl() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	elems := doc.Find("div")
	elems.Each(func(i int, selection *goquery.Selection) {
		println(selection.Text())
	})
}

func NodesFromFile(fl string) {
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
	tb := doc.Find("#tab1 tbody")

	//divs.Each(func(i int, sel *goquery.Selection) {
	//fmt.Sprintf("Node %s :: %d :: %s", sel.Nodes[0].Data, sel.Length(), sel.Text())
	//println("====")
	//println(sel.Text())
	//})

	//tb.Children().Each(func(i int, el *goquery.Selection) {
	//	fmt.Println(el.Html())
	//})

	//trs := doc.Find("tr")
	//fmt.Println(trs.Text())

	//tb.Each(func(i int, sl *goquery.Selection) {
	//	println(sl.Text())
	//})

	tb.Find("tr").Each(func(i int, el *goquery.Selection) {
		fmt.Println(el.Text())
	})
}
