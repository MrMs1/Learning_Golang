package mylib

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Goroutine() {
	var wg sync.WaitGroup
	wg.Add(1) // 並列処理の数をセット
	go goroutine("world", &wg)
	normal("hello")
	//time.Sleep(1000 * time.Millisecond)
	wg.Wait() //Doneが呼ばれるまで待つ
}
