package main

import (
	"fmt"
	"strconv"
	// leetcode "leetcode-dsa-grind/leetcode/two_sum"
)

func main() {
	// fmt.Println(leetcode.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(IsPalindrome(12121))
}

func IsPalindrome(x int) bool {
	//cast to string
	stringInt := strconv.Itoa(x)
	halfLenArray := len(stringInt)/2
	lastIndexArray := len(stringInt)
	//loop the string 
	for i := 0; i < halfLenArray; i++ {
		if stringInt[i] != stringInt[lastIndexArray-i-1] {
			return false
		}
	}
	return true
}