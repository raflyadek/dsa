package main

import (
	"leetcode-dsa-grind/week2/dsa"
)

func main() {

	newNode := dsa.NewNode(5)
	linkedList := dsa.NewLinkedList(newNode)

	linkedList.Insert(3)
	linkedList.Insert(2)
	linkedList.Insert(1)
	linkedList.Insert(123)

	// linkedList.Delete()

	linkedList.Print()
}