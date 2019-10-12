package main

import (
	"encoding/json"
	"fmt"
)

type configuration struct {
	Data []byte `json:"data"`
}

func main() {
	// Use the byte array/slice data type in your JSON object, but the binary data will have to be base64 encoded.
	raw := []byte(`{"data":"wg=="}`)
	var decoded configuration

	if err := json.Unmarshal(raw, &decoded); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v", decoded) //prints: main.config{Data:[]uint8{0xc2}}
}
