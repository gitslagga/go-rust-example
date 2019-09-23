package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

type admin struct {
	user
	level string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

func main() {
	ad := admin{
		user: user{
			name: "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}
	sendNotification(&ad)
}

func sendNotification (n notifier) {
	n.notify()
}
