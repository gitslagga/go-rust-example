package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var message = make(chan string)

func WriteMsgToClient(clnt Client, conn net.Conn) {
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + ": " + msg
	return
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	hasData := make(chan bool)

	netAddr := conn.RemoteAddr().String()
	clnt := Client{make(chan string), netAddr, netAddr}
	onlineMap[netAddr] = clnt
	go WriteMsgToClient(clnt, conn)

	message <- MakeMsg(clnt, "login")
	isQuit := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("检测到客户端：%s退出\n", clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err", err)
				return
			}
			msg := string(buf[:n-1])
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				for _, user := range onlineMap {
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				clnt.Name = newName
				onlineMap[netAddr] = clnt
				conn.Write([]byte("rename successful\n"))
			} else {
				message <- MakeMsg(clnt, msg)
			}
			hasData <- true
		}
	}()

	for {
		select {
		case <-isQuit:
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout")
			return
		case <-hasData:
			// do nothing, reset timer
		case <-time.After(time.Second * 60):
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "time out leaved")
			return
		}
	}
}

func Manager() {
	onlineMap = make(map[string]Client)

	for {
		msg := <-message
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("net.Listen err", err)
		return
	}

	defer listener.Close()

	go Manager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
			return
		}

		go HandlerConnect(conn)
	}
}
