package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		asArray := Permutate(text)
		sort.Strings(asArray)
		for index, value := range asArray {
			if index == len(asArray)-1 {
				fmt.Printf("%s\n", value)
			} else {
				fmt.Printf("%s,", value)
			}
		}
	}
}

// [abc] => [abc, bac, bca, acb, cab, cba]
func Permutate(text string) (strArray []string) {
	if len(text) == 1 {
		strArray = append(strArray, text)
		return
	}

	permArray := Permutate(text[1:])
	head := text[0] //type char
	for _, value := range permArray {
		i := 0
		for i <= len(value) {
			//			fmt.Printf("pre: %s\n",value[:i])
			//			fmt.Printf("head: %s\n",string(head))
			//			fmt.Printf("post: %s\n", value[i:])
			newString := strings.Join([]string{value[:i], string(head), value[i:]}, "")
			strArray = append(strArray, newString)
			i++
		}
	}
	return
}

//func PowInt(base int, power int) (result int) {
//	if power == 0  {
//		result = 1
//		return
//	}
//	result = base
//	power--
//	for power > 0 {
//		result *= base
//		power--
//	}
//	return
//}

//func Permutate(text string) {
//	// digits < upper case < lower case
//	// 48-57 digits < 65-90 upper case < 97-122 lower case
//	asString := strings.Split(text, "")
//	sort.Strings(asString)
//	i := 1
//	for i <= Factorial(len(asString)) {
//		InOrder(asString, i)
//		i++
//	}
//}
//
//func InOrder(array []string, order int) {
//	if order == Factorial(len(array)) {
//		sort.Reverse(array)
//		fmt.Println(strings.Join(array,""))
//	} else if order == 1{
//		fmt.Printf("%s,",strings.Join(array,""))
//	} else {
//
//	}
//}

//func Factorial(len int) (out int) {
//	if len == 1 {
//		out = 1
//		return
//	} else {
//		out = len * Factorial(len-1)
//		return
//	}
//}
