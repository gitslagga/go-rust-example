package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	records := [][]byte{
		[]byte(`{"status": 200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			fmt.Println("error:", err)
			return
		}

		var sstatus string
		if err := json.Unmarshal(result.Status, &sstatus); err == nil {
			result.StatusName = sstatus
		}

		var nstatus uint64
		if err := json.Unmarshal(result.Status, &nstatus); err == nil {
			result.StatusCode = nstatus
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}

// [0] result => {StatusCode:200 StatusName: Status:[50 48 48] Tag:one}
// [1] result => {StatusCode:0 StatusName:ok Status:[34 111 107 34] Tag:two}
