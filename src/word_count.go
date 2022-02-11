package main

import "fmt"

func main() {
	var words = []string{"foo", "foo", "foo", "bar", "bar", "baz"}
	var result = map[string]int32{}
	for _, w := range words {
		result[w] += 1
	}
	fmt.Println("The word count is", result)
}
