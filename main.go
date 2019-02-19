package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.mozilla.org/en-US/firefox/releases/")
	if err != nil {
		fmt.Println("failed to scrape url")
	}
	doc.Find("html").Each(func(_ int, s *goquery.Selection) {
		version, _ := s.Attr("data-latest-firefox")
		fmt.Println("latest firefox is", version)
	})
}
