package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Closing HTTP Connections
	req, err := http.NewRequest("GET", "https://api.ipify.org?format=json", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Close = true
	//or do this:
	//req.Header.Add("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(string(body)))
}
