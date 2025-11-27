package main

import "fmt"

func isPalindrom(s string) bool {
	length := len(s)
	//rune is a way in go to turn string to a array of string kinda like put the string int the array so we can indexing it
	char := string(s)

	for i := 0; i < length/2; i++ {
		if char[i] != char[length-i-1] {
			return false
		}
	}
	return true
}


func main() {
	// str := "hello"
	// fmt.Println(str[1])
	fmt.Println(isPalindrom("katak"))
}