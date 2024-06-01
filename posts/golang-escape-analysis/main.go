package main

import "fmt"

func main() {
	dat := make(map[string]interface{})
	dat["BlogName"] = "Kyden's Blog"
	val := dat["BlogName"]
	fmt.Printf("%v\n", val)
}
