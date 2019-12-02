package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Network listen socket")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Listen client connection")

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	conn.Write(buf[:n])
	fmt.Println("Server Read Data:", string(buf[:n]))
}
