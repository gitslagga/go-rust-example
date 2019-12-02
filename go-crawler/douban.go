package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	url := "https://movie.douban.com/top250"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Chrome")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find(".hd a span").Eq(0).Text()
		quote := s.Find(".quote .inq").Text()
		fmt.Printf("Review %d: %s - %s\n", i, title, quote)
	})
}

func main() {
	ExampleScrape()
}
