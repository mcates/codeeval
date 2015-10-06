package main

import (
	"fmt"
	"strconv"
)

func main() {
	max := 1000
	save_max := 1
	for i := 1; i < max; i++ {
		if IsPalindromeAndPrime(i) {
			save_max = i
		}
	}
	fmt.Println(save_max)
}

func IsPalindromeAndPrime(val int) (ret bool) {
	valAsString := strconv.Itoa(val)
	valAsRuneArray := []rune(valAsString)
	varLength := len(valAsString)
	ret = true
	if varLength%2 == 0 {
		for i := 0; i < varLength/2; i++ {
			if valAsRuneArray[i] != valAsRuneArray[varLength-i-1] {
				ret = false
				return
			}
		}
	} else {
		for i := 0; i < varLength+1/2; i++ {
			if valAsRuneArray[i] != valAsRuneArray[varLength-i-1] {
				ret = false
				return
			}
		}
	}
	if val%2 == 0 {
		for i := 2; i < val/2; i++ {
			if val%i == 0 {
				ret = false
				return
			}
		}
	} else {
		for i := 2; i < (val+1)/2; i++ {
			if val%i == 0 {
				ret = false
				return
			}
		}
	}

	return
}
