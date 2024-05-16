package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	fmt.Println("Hello, World!")

	wg.Add(1)
	go f(&wg)

	wg.Wait()
}

func f(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("goroutine")
}
