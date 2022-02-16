package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	head := &Node{
		Value: 0,
		Next:  nil,
	}

	node := head
	for _, num := range nums {
		node.Next = &Node{
			Value: num,
			Next:  nil,
		}
		node = node.Next

	}

	node = head
	for node.Next != nil {
		node = node.Next
		fmt.Printf("%d, ", node.Value)
	}
}
