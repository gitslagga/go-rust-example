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
	name string
	email string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s\n", a.name, a.email)
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

func sendNotification (n notifier) {
	n.notify()
}
