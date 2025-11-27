package main

import "fmt"

func countCharacters(s string) map[rune]int {
	count := make(map[rune]int)
	isFound := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if !isFound[string(char)] {
			isFound[string(char)] = true
			count[rune(char)] = 1
		} else {
			count[rune(char)]++
		}
	}
	return count
}

func main() {
	fmt.Println(countCharacters("abcdeabc"))
}