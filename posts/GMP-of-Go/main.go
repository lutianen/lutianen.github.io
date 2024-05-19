package main

import (
	"fmt"
	"sync"
	"time"
)

var mtx sync.Mutex

func main() {
	go func() {
		mtx.Lock()
		defer mtx.Unlock()

		fmt.Printf("Start\n")
		time.Sleep(time.Second * 10)
		fmt.Printf("End\n")
	}()

	time.Sleep(time.Second) // Ensure child goroutine gets the mutex before main goroutine

	fmt.Printf("Try to acquire mutex\n")
	mtx.Lock()
	fmt.Printf("Main goroutine\n")
	mtx.Unlock()
}
