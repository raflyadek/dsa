package main

import (
	"fmt"
	"strings"
)

func countVocal(s string) int {
	//count vocal
	result := 0
	vocal := "aiueo"
	//for loop
	// for i := 0; i < len(s); i++ {
	// 	for j := 0; j < len(vocal); j++ {
	// 		if s[i] == vocal[j] {
	// 			result++
	// 		}
	// 	}

	//for range
	for _, v := range s {
		fmt.Println(string(v))
		if strings.Contains(vocal, string(v)) {
			result++
		}
		fmt.Println("---")
	}

	return result
	}

	//anagram
	// if len(s1) != len(s2) {
	// 	return false
	// }
	// count := 0
	// for _, str := range s1 {
	// 	if strings.Contains(s2, string(str)) {
	// 		count++
	// 	}
	// }
	// if count == len(s1) {
	// 	return true
	// }

	// factorial
	// if n == 1 {
	// 	return 1
	// }
	// return n * countVocal(n-1)
	// return false
	// return false
// }

func main() {
	fmt.Println(countVocal("aiueo"))
	// fmt.Println(strings.Split("abdcds", ""))
}