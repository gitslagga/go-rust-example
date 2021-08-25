package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GZipDecompress(input []byte) (string, error) {
	buf := bytes.NewBuffer(input)
	reader, gzipErr := gzip.NewReader(buf)
	if gzipErr != nil {
		return "", gzipErr
	}
	defer reader.Close()

	result, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return "", readErr
	}

	return string(result), nil
}

func GZipCompress(input string) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write([]byte(input))
	if err != nil {
		return nil, err
	}

	err = gz.Flush()
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type PingMessage struct {
	Ping int64 `json:"ping"`
}

func ParsePingMessage(message string) *PingMessage {
	result := PingMessage{}
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		return nil
	}

	return &result
}

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)

	urlTarget := fmt.Sprintf("wss://api.huobi.pro/ws")
	purl, err := url.Parse("http://127.0.0.1:1080")
	if err != nil {
		log.Fatal(err)
	}

	dialer := websocket.Dialer{
		Proxy:            http.ProxyURL(purl),
		HandshakeTimeout: 45 * time.Second,
	}

	var conn *websocket.Conn
	conn, _, err = dialer.Dial(urlTarget, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	symbol := "btcusdt"
	topic := fmt.Sprintf("market.%s.bbo", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\", \"id\": \"%s\"}", topic, "5")

	err = conn.WriteMessage(websocket.TextMessage, []byte(sub))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case <-quit:
				return

			default:
				if conn == nil {
					return
				}

				msgType, buf, err := conn.ReadMessage()
				if err != nil {
					log.Printf("conn ReadMessage err: %v\n", err)
					return
				}

				if msgType == websocket.BinaryMessage {
					message, err := GZipDecompress(buf)
					if err != nil {
						log.Printf("buf GZipDecompress err: %v\n", err)
						return
					}

					pingMsg := ParsePingMessage(message)
					if pingMsg != nil && pingMsg.Ping != 0 {
						log.Printf("Received Ping: %d\n", pingMsg.Ping)

						pongMsg := fmt.Sprintf("{\"pong\": %d}", pingMsg.Ping)
						err := conn.WriteMessage(websocket.TextMessage, []byte(pongMsg))
						if err != nil {
							log.Printf("conn WriteMessage err: %v\n", err)
							return
						}

						log.Printf("Replied Pong: %d\n", pingMsg.Ping)
					} else {
						log.Println(message)
					}
				}
			}
		}
	}()

	sg := <-quit
	fmt.Printf("receive the signal:%v\n", sg)
}
