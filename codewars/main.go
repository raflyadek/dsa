package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(NumberToString(25))
	// fmt.Println(CheckForFactor(25, 5))
	// fmt.Println(GetGrade(30, 40, 50))
	// fmt.Println(Past(0, 0, 1))
	// fmt.Println(TwoSort([]string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}))
	// fmt.Println(Accum("abcde"))
	// fmt.Println(ReverseWords("rafly   ade KusumA"))
	// fmt.Println(AbbrevName("rafly ade"))
	// fmt.Println(IsValidWalk([]rune{'n','s','n','s','n','s','n','s','n','s'}))
	// fmt.Println(GetCount("abcde"))
	// fmt.Println(GetMiddle("raflya"))
	// fmt.Println(HighAndLow("6 7 8 9 11 2"))
	// fmt.Println(FindShort("rafly ade kusuma aa ii uu"))
	// fmt.Println(FindUniq([]float32{1, 1, 1, 1, 0.55, 1}))
	// fmt.Println(StringToNumber("1234"))
	// fmt.Println(DNAtoRNA("GCAT"))
	// fmt.Println(Is_valid_ip("127.20.32.84"))
	// fmt.Println(Score([5]int{5, 5, 5, 2, 1}))
	// var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	// var a2 = []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
	// fmt.Print(Comp(a1, a2))
	numbers := []int{9, 7, 5, 4, 3, 2}
	fmt.Println(BubblesortOnce(numbers))
}

func FindMultiples(integer, limit int) []int {
  // Your code here!
if integer < 1 {
  return nil
}
  
  if limit < integer {
    return nil
  }

//   length := limit/integer
  result := make([]int, 0)
  current := integer
  for current <= limit {
	result = append(result, current)
	current += integer
  }
  return result
}

func RemoveChar(word string) string {
	// if word == "" {
	// 	return ""
	// }
	// result := ""
	// // lastIndex := len(word) - 1
	// // firstIndex := 0

	// length := len(word)
	// for i := 1; i < length-1; i++ {
	// 	//if words index 0 or index last continue or not add to result
	// 	//silly because just start after the first index and stop before the last index
	// 	result += string(word[i])
	// }

	// return result

	//solutin 2
	return word[1:len(word)-1]
}	

// func FindMultiples(integer, limit int) []int {
//   current := integer
//   nums := make([]int, 0)
  
//   for current <= limit {
//     nums = append(nums, current)
//     current += integer
//   }
//   return nums
// }

func CountPositivesSumNegatives(numbers []int) []int {
	resultPositive := 0
	resultNegative := 0
	for _, v := range numbers {
		//silly solution because the positive is count not sum the negative is sum
		// resultPositive += v
		// if v < 0 {
		// 	resultNegative += v
		// }
		if v > 0 {
			resultPositive++
		} else {
			resultNegative += v
		}
	}
	var res []int
	res = append(res, resultPositive, resultNegative)
	return res // your code here
}

func IsPalindrome(str string) bool {
	if str == "" {
		return false
	}
	s := strings.TrimSpace(str)
	s = strings.ToLower(s)
	for i := 0; i < len(str)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
//bin to decimal is given [1 1 1 1 1] so the first index binary is 1 the next is 2 and 4 and 8 and 16 and 32 and 64 and etc.. 
//OR 2^ <- and the lets say its first index/0 <- so its 2^0 = 1 and 2^1 = 2 etc
//you got the point
//so the string when we access it via indexing it returns a byte, and that byte lets say 32 
//32 <- so there are 2 binary and we just multiply it with the binary like the first index is 2^0 * 2 = 2 and 2^1 * 3 = 6 so 6+2 = 8
//above is false because binary only 0 and 1 so if its 11 == 2^0 * 1 = 1 + 2^1 * 2 = 3 == 3+1 = 4
func BinToDec(bin string) int {
	result := 0
	for i, _ := range bin {
		//start fro last index because binary start with the last index 
		fromLast := bin[len(bin)-i-1]
		fromLastStr := string(fromLast)
		//convert the input to int
		fromLastInt, err := strconv.Atoi(fromLastStr)
		if err != nil {
			return 0
		}
		//2 power i and then * int from input
		fromLastIntToDec := (int(math.Pow(2, float64(i)))) * fromLastInt
		//sum the decimal to the result 
		result += fromLastIntToDec
	}

	//solution 2
	str, _ := strconv.ParseInt(bin, 2, 64)

	return int(str)
}

func NearestSq(n int) int {
	// Code goes here
	//square from parameter
	// result := math.Sqrt(float64(n))
	// //floor the result to below
	// resultFloor := math.Floor(result)
	// //get the decimal
	// _, decimal := math.Modf(result)

	// //for return if below then just the resultfloor but if upper then + 1
	// below := int(resultFloor * resultFloor)
	// upper := int((resultFloor + 1) * (resultFloor + 1))  
	// //if decimal < 0.5 then do
	// if math.Abs(decimal) < 0.5 {
	// 	return below
	// } else {
	// 	return upper
	// }

	//solution 2 yeah its so much better than floor it is round it the nearest integer 
	result := math.Round(math.Sqrt(float64(n)))
	//if the decimal is below .5 then resultFloor else resultFloor+1
	return int(result * result)
}


func HowMuchILoveYou(i int) string {
  	if i < 0 {
  		return ""
	}
	result := [6]string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}
	// length := len(result)
  	return result[(i - 1) % 6]
}

func NumberToString(n int) string {
	res := strconv.Itoa(n)
	return res
}

func CheckForFactor(base int, factor int) bool {
    if base % factor != 0 {
		return false
	}

	return true
}

func GetGrade(a,b,c int) rune {
	averageGrade := (a + b + c) / 3
	fmt.Println(averageGrade)

    switch {
	case averageGrade >= 90:
		return 'A'
	case averageGrade >= 80:
		return 'B'
	case averageGrade >= 70:
		return 'C'
	case averageGrade >= 60:
		return 'D'
	default:
		return 'F'
	}
}

func Past(h, m, s int) int {
	// s = s * 1000
	// m = m * 60000
	// h = h * 360000
	// result := s + m + h
	// return result

	//solution 2 leverage the return 
	return (h*3600000 + m*60000 + s*1000)
}

func TwoSort(arr []string) string {
	result := ""
	
	// slices.Sort(arr)
	// sort := arr[0]
	// for _, v := range sort {
	// 	result += string(v) + "***"
	// }

	//solution 2
	sort.Strings(arr)
	sortt := arr[0]
	for i, v := range sortt {
		if i == len(sortt)-1 {
			result += string(v)
		} else {
			result += string(v) + "***"	
		}
	}
	return result
}

func Accum(s string) string {
    // your code
	result := ""
	//so if there is inner loop and outer loop / 2loop
	//the i is already increment after it loop 
	for i := 0; i < len(s); i++ {
		result += strings.ToUpper(string(s[i]))
		for j := 0; j < i ; j++ {
			result += strings.ToLower(string(s[i]))
		}
		if i != len(s)-1 {
			result += "-"
		}
	}

	return result
}

func ReverseWords(str string) string {
	runes := []rune(str)
	result := ""
	words := ""
	
	for i := 0; i < len(runes); i++ {
		if runes[i] != ' ' {
			words += string(runes[i])
		} else {
			for j := len(words)-1; j >= 0; j-- {
				result += string(words[j])
			}
			//empty the words
			words = ""
			result += " "
		}
	}

	if len(words) > 0 {
		for i := len(words)-1; i >= 0; i-- {
			result += string(words[i])
		}
	}
	return result

	// words := strings.Fields(str)
	// result := []string{}
	
	// for _, v := range words {
	// 	runes := []rune(v)
	// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	// 		runes[i], runes[j] = runes[j], runes[i]
	// 	}
	// 	result = append(result, string(runes))
	// }

	// return strings.Join(result, " ")
}

func AbbrevName(name string) string{
	result := ""
	//your code here
	//Rafly Ade  -> R.A
	word := strings.Fields(name)
	for i, v := range word {
		result += strings.ToUpper(string(v[0]))
		if i != len(word)-1 {
			result += "."
		}
	}
	return result
}

func IsValidWalk(walk []rune) bool {
	//each block is 1 minute so the walk rune is always 10
	if len(walk) != 10 {
		return false
	}
	
	//each walk is a direction in compass 'n' 's' 'e' 'w'
	if walk[0] == 'e' && walk[len(walk)-1] == 'w' {
		return true
	} else if walk[0] == 'n' && walk[len(walk)-1] == 's' {
		return true
	} else {
		return false
	}
}

func GetCount(str string) (count int) {
	vowel := "aiueo"
	for _, v := range str {
		if strings.Contains(vowel, string(v)) {
			count++
		}
	}
	return count
}

func GetMiddle(s string) string {
	//Code goes here!
	//middle if odd is 2 else is 1
	middle := len(s)/2
	if len(s) % 2 == 0 {
		return s[middle-1:middle+1]
	} else {
		return string(s[middle])
	}
}

func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func HighAndLow(in string) string {
    fields := strings.Fields(in)
    if len(fields) == 0 {
        return ""
    }
	result := ""

    // Convert the first number
    num, err := strconv.Atoi(fields[0])
    if err != nil {
        return ""
    }

    high, low := num, num

    for _, v := range fields {
        n, err := strconv.Atoi(v)
        if err != nil {
            continue // ignore invalid entries
        }

        if n > high {
            high = n
        }
        if n < low {
            low = n
        }
    }
	highStr := strconv.Itoa(high)
	lowStr := strconv.Itoa(low)
	result += highStr +" "+ lowStr
    return result
}

func FindShort(s string) int {
	//your code

	//make input into slice of string 
	fields := strings.Fields(s)
	sort.Strings(fields)
	//set result to first index length of fields 
	result := len(fields[0])

	for _, v := range fields {
		if result > len(v) {
			result = len(v)
		}
	}
	return result
}

func FindUniq(arr []float32) float32 {
	// Do the magic
	if len(arr) < 3 {
		return 0
	}

	//map
	result := make(map[float32]int)
	for i := 0; i <len(arr); i++ {
		result[arr[i]]++
	}
	
	for i, v := range result {
		if v == 1 {
			//return the key if value is 1
			return float32(i)
		}
	}
	return 0
}

func StringToNumber(str string) int {
	// your code here
	result, _ := strconv.Atoi(str)

	return result
}

func DNAtoRNA(dna string) string {
	// your code here
	//every t replaced by u
	return strings.ReplaceAll(dna, "T", "U")
	// result := ""

	// for _, v := range dna {
	// 	if string(v) != "T" {
	// 		result += string(v)
	// 	} else {
	// 		result += "U"
	// 	}
	// }
	// return result
}

func Is_valid_ip(ip string) bool {
	//split the ip into numerik only 
	ipAddr := strings.Split(ip, ".")
	
	for _, v := range ipAddr {
		//casting the ip to int 
		ipInt, err := strconv.Atoi(v)
		if err != nil {
			return false
		}

		//check if the ip is in constraints
		if ipInt > 255 || ipInt < 0 {
			return false
		}
	}
	
	return true
}

func Score(dice [5]int) int {
	result := 0
	scoreMap := make(map[int]int)
	for _, v := range dice {
		scoreMap[v]++
	}
	
	for i, v := range scoreMap {
		//if v >= 3 and after that v - 3 because if there are leftovers it still count 
		if i == 1 && v >= 3 {
			result += 1000
			v = v - 3
		}
		//if v >= 3 and after that v - 3 because if there are leftovers it still count 
		if v >= 3 {
			result += i * 100 
			v = v - 3
		}
		if i == 1 && v < 3 {
			result += v * 100
		}
		if i == 5 && v < 3 {
			result += v * 50
		}

	}
	return result
}

func Comp(array1 []int, array2 []int) bool {
	//basic check
	if len(array1) != len(array2) {
		return false
	}

	if array1 == nil || array2 == nil {
		return false
	}

    // your code
	//square array1
	squares1 := make([]int, len(array1))
	for i, v := range array1 {
		squares1[i] = v * v
	}

	//sort both array
	sort.Ints(squares1)
	sort.Ints(array2)
	fmt.Println(squares1)
	fmt.Println(array2)
	for i := range squares1 {
		if squares1[i] != array2[i] {
			return false
		}
	}
	return true
}

func SumDigPow(a, b uint64) []uint64 {
	//a & b is range
	return nil
}

func Rps(p1, p2 string) string {
	if len(p1) == len(p2) {
		return "Draw!"
	}	
	if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	} else if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	} else if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	}

	return "Player 2 won!"
}

func BubblesortOnce(numbers []int) []int {
	n := len(numbers)
	//compare the first index to second and move it untill it reaches end
	for i := 0; i < n-1; i++ {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
		}
	}
	return numbers
}
