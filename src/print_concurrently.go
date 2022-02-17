package main

import "time"

func main() {

	var arr = make([]int, 10)
	for i := range arr {
		arr[i] = i
	}
	for _, num := range arr {
		go println(num)
	}

	time.Sleep(3 * time.Second)
}
