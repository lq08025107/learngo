package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("sending user email to %s to <%s>\n", u.name, u.email)
}
func main() {
	u := user{"Bill", "bill@domain.com"}
	sendNotification(u)
}

func sendNotification(n notifier) {
	n.notify()
}
