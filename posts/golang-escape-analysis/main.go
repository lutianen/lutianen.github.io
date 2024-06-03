package main

import "fmt"

func f(a *int) {
	fmt.Printf("%v\n", *a)
	return
}

func main() {
	a := 10
	fn := f
	fn(&a)
	fmt.Printf("a = %v\n", a)
}
