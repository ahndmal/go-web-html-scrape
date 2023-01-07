package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	NodesFromFile("/home/andrii/GolandProjects/go-web-html-scrape/index.html", "div")

	fmt.Println(fmt.Sprintf(" Time taken: %d", time.Now().Sub(start).Milliseconds()))
}

func FromUrl() {
	url2 := "https://en.wikipedia.org/wiki/Project_Waler"
	//url1 := "https://en.wikiversity.org/wiki/Wikiversity:Main_Page"
	resp, err := http.Get(url2)
	if err != nil {
		fmt.Println(err)
	}
	texts := make([]string, 0)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	elems := doc.Find("div")
	elems.Each(func(i int, selection *goquery.Selection) {
		texts = append(texts, strings.TrimSpace(selection.Text())) //selection.Text()
	})
	fmt.Println(len(texts))
	for _, text := range texts {
		fmt.Println(text)
	}
}

func NodesFromFile(fl string, toFind string) {
	file, err := os.ReadFile(fl)
	if err != nil {
		fmt.Println(err)
	}
	reader := bytes.NewReader(file)
	//fs.ReadFile(, "index.html")
	//io.ReadAll(fs.File)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
	}
	divs := doc.Find(toFind)
	divs.Each(func(i int, sel *goquery.Selection) {
		//fmt.Println(sel.Text())
		fmt.Printf("Node %s :: %d :: %s", sel.Nodes[0].Data, sel.Length(), sel.Text())
	})
}
