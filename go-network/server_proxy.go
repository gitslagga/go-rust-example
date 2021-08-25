package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func httpUserInfo(w http.ResponseWriter, r *http.Request) {
	ret := "your addr is:" + r.RemoteAddr + "\r\n"
	ret += "request headers:" + fmt.Sprint(r.Header)

	io.WriteString(w, ret)
}

func main() {
	targetURL, err := url.Parse("http://127.0.0.1:80")
	log.Println("program start......")
	if err != nil {
		log.Fatal("bad url:", err)
	}

	httpProxy := httputil.NewSingleHostReverseProxy(targetURL)
	http.HandleFunc("/userInfo/", httpUserInfo)
	http.Handle("/", httpProxy)

	log.Println("listen http proxy on :8001.....")
	http.ListenAndServe(":8001", nil)
}
