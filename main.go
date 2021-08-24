package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var STARTER_URL = "https://en.wikipedia.org/wiki/Puerto_Rico"

func main() {
	res, err := http.Get(STARTER_URL)

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div#mw-content-text a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "https") {
				fmt.Println(text, href)
			}
		}
	})
}
