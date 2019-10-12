package main

import (
	"encoding/json"
	"fmt"
)

type config struct {
	Data string `json:"data"`
}

func main() {
	// JSON String Values Will Not Be Ok with Hex or Other non-UTF8 Escape Sequences
	raw := []byte(`{"data":"\\xc2"}`)

	var decoded config

	json.Unmarshal(raw, &decoded)

	fmt.Printf("%#v", decoded) //prints: main.config{Data:"\\xc2"}
	//todo: do your own hex escape decoding for decoded.Data
}
