package main

import (
	"fmt"
	leetcode "leetcode-dsa-grind/leetcode/two_sum"
)

func main() {
	fmt.Println(leetcode.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(IsPalindrome(1241))
}

func IsPalindrome(x int) bool {
	intArray := make([]int, 0, 0)
	intArray = append(intArray, x)
	halfLenArray := len(intArray)/2
	lastIndexArray := len(intArray)
	for i := 0; i < halfLenArray; i++ {
		if intArray[i] != intArray[lastIndexArray-i-1] {
			return false
		}
	}
	return true
}