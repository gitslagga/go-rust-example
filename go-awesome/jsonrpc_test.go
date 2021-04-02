package go_awesome

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

type UserService struct {
}
type User struct {
	UserId int
	Name   string
}

func (UserService) GetUser(userId int, result *User) error {
	result.UserId = userId
	result.Name = "Jack"
	return nil
}

func TestRpc(t *testing.T) {
	rpc.Register(UserService{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	rpc.ServeConn(conn)
}

func TestRpcClient(t *testing.T) {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	user := User{}
	err = client.Call("UserService.GetUser", 1, &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)
}

func TestJsonRpc(t *testing.T) {
	rpc.Register(UserService{})
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go jsonrpc.ServeConn(conn)
	}
}

func TestJsonRpcClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	user := User{}
	err = client.Call("UserService.GetUser", 1, &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)
}
