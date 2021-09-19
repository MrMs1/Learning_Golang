package mylib

import "fmt"

func goroutine3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func Channellesson2() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine3(s, c)
	for i := range c {
		fmt.Println(i)
	}
}
