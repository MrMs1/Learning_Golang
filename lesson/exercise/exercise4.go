package exercise

import (
	"fmt"
)

func goroutineex(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func Exercise4() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutineex(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
