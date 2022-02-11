package main

import "fmt"

func main() {
	var result int32 = 0
	mapping := map[string]int32{
		"foo": 1, "bar": 1, "baz": 3}
	for _, val := range mapping {
		result += val
	}
	fmt.Println("Sum the values of mapping is", result)
}
