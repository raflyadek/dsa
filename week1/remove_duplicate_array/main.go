package main

import "fmt"

func removeDuplicates(arr []int) []int {
	//create map to be the flag 
	isFound := make(map[int]bool)
	result := []int{}
	for i := 0; i < len(arr); i++ {
		//check if isfound false then add it to the result
		if !isFound[arr[i]] {
			//set the isFound with key input[i]true so it wont getting to this if condition again
			isFound[arr[i]] = true
			//append the input with i index to the result variable
			result = append(result, arr[i])
		}
	}
	return result
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 3, 3, 4, 5, 5, 6}))
}
