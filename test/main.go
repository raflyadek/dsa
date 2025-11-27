package main

import "fmt"

func main() {
	//bubble sort
	numbers := []int{9, 5, 7, 2, 1}
    fmt.Println("Before sorting:", numbers)
    bubbleSort(numbers)
    fmt.Println("After sorting:", numbers)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}