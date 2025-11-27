package main

import (
	"fmt"
)

func reverseString(s string) string {
	if s == "" {
		return ""
	}

	// //variable to store array string and then join it to convert it to string
	// arrayString := []string{}
	
	// //so with string we can just straight to indexing 
	// //use for loop to iterate every index of input
	// for i := 0; i < len(s); i++ {
	// 	//if the length of string is 5 so if we substract it with i - 1 it will be 4 whicih the last index 
	// 	//because index start at 0 
	// 	reverse := s[len(s)-i-1]
	// 	//because indexing string returns a byte we have to casting it to string
	// 	reverseString := string(reverse)
	// 	//append to arraystring
	// 	arrayString = append(arrayString, reverseString)
	// }
	// //casting the slice of string to string 
	// //variable to hold the reverse string
	// reversedString := strings.Join(arrayString, "")
	// length := len(s)
	// reversedString := make([]rune, length)

	// for _, v := range s {
	// 	length--
	// 	reversedString[length] = v
	// 	log.Println(reversedString)
	// }
	
	// return string(reversedString[length:])
	//beatufiul, for every string(v) it gonna replaced in the index 0
	result := ""
	for _, v := range s {
		//r
		//ar
		//far
		//lfarcd 
		//ylfar
		result = string(v) + result
	}
	return result
}

func main() {
	// log.Println(reverseString("abcdef"))
	// log.Println(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
	// log.Println(IsPalindrome("madaM"))
	// log.Println(BinToDec("11"))
	fmt.Println(reverseString("rafly"))
}