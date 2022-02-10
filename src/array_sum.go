package main

import "fmt"

func main() {
	var num = []int{1, 2, 3, 4}
	result := 0
	for _, val := range num {
		result += val
	}
	fmt.Println("Sum(num) is", result)
}
