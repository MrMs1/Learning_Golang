package mylib

import (
	"fmt"
	"time"
)

func goroutinecs1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutinecs2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func ChannelSelectlesson() {
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutinecs1(c1)
	go goroutinecs2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
