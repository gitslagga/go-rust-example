package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func useProxy() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:7890")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}

	resp, err := client.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func main() {
	useProxy()
}
