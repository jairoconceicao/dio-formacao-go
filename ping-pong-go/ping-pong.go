package main

import (
	"fmt"
	"time"
)

func Ping(c chan string) {
	for {
		c <- "Ping"
		time.Sleep(100 * time.Millisecond)
	}
}

func Pong(c chan string) {
	for {
		c <- "Pong"
		time.Sleep(100 * time.Millisecond)
	}
}

func Print(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan string)

	go Ping(ch)
	go Print(ch)
	go Pong(ch)

	var chKey string
	fmt.Scanln(&chKey)
}
