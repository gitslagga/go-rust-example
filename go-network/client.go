package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		fmt.Println("net.Dail err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Network dial connection")

	n, err := conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	buf := make([]byte, 4096)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("Client read data:", string(buf[:n]))
}
