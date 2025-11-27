package main

import "fmt"

func reversedArray(s []string) []string {
	//variable to store slice of string
	sliceOfString := []string{}
	
	//using for range?
	for i := 0; i < len(s); i++ {
		//okay so if we indexing the array we get the type not byte 
		reverse := s[len(s)-i-1]
		sliceOfString = append(sliceOfString, reverse)
	}

	return sliceOfString
}

func main() {
	fmt.Println(reversedArray([]string{"apple", "banana", "orange"}))
}