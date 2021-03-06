package thirdpartylib

import (
	"context"
	"fmt"
	"lesson2/mylib"
	"time"

	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Cound not get lock")
		return
	}

	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func Semaphore() {
	mylib.StartLog()
	defer mylib.EndLog()

	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
}
