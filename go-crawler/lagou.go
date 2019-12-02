package main

import (
	"fmt"

	"github.com/eddieivan01/nic"
)

func main() {

	// session 保持Cookie
	session := &nic.Session{}
	session.Get("https://www.lagou.com/jobs/list_go?labelWords=&fromSearch=true&suginput=", &nic.H{
		Headers: nic.KV{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36",
		},
	})

	resp, _ := session.Post("https://www.lagou.com/jobs/positionAjax.json", &nic.H{
		Data: nic.KV{
			"first":               "true",
			"pn":                  "1",
			"kd":                  "go",
			"city":                "深圳",
			"needAddtionalResult": "false",
		},
		Headers: nic.KV{
			"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36", //     "referer": "https://www.lagou.com/jobs/list_go?city=%E6%B7%B1%E5%9C%B3&cl=false&fromSearch=true&labelWords=&suginput=&labelWords=hot",
			"Referer":      "https://www.lagou.com/jobs/list_go?city=%E6%B7%B1%E5%9C%B3&cl=false&fromSearch=true&labelWords=&suginput=&labelWords=hot",
			"Origin":       "https://www.lagou.com",
			"Accept":       "application/json, text/javascript, */*; q=0.01",
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		},
	})
	fmt.Println(resp.Text)
}
