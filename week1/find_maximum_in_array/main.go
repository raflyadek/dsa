package main

import "fmt"

func findMax(arr []int) int {
	result := arr[0]

	// for i := 0; i < len(arr); i++ {
	// 	lengthArr := len(arr)
	// 	// fmt.Println(result)
	// 	resultNext := arr[lengthArr-i-1]
	// 	if resultNext > result {
	// 		result = resultNext
	// 	}
	// 	// fmt.Println(result)
	// }

	//use for range 
	for _, max := range arr {
		if max > result {
			result = max
		}
	}
	return result
}

func main() {
	fmt.Println(findMax([]int{1, 2, 3, 4, 100, 2, 4, 9}))
}
