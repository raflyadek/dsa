package main

import (
	"fmt"
	"log"
)

// import "fmt"

func sumOfDigits(n int) int {
	//check n
	if n == 0 {
		return 0
	}

	//so if the input is 1234 we want 1+2+3+4 = 10
	//create variable to store the value
	result := 0

	//create rune to store the digits in byte? 
	// digit := []byte(n)

	//for range to indexing the input int?
	// for _, digit := range n {
	// 	fmt.Println(digit)
	// }

	//using for statement to just brute force it?
	for n > 0 {
		//using mod 10 to get every last digit 
		result += n % 10
		log.Print(result)
		//divide n with 10 to get rid of the last digit
		n = n / 10
	}

	return result
}

func main() {
	fmt.Println(sumOfDigits(1234))
}