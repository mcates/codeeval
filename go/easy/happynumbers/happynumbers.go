package main

import (
	"os"
	"errors"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("Include file name as first parameter"))
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		b := HappyNumbers(line)
		fmt.Println(b)
	}
}

func HappyNumbers(s string) int {
	sum := 2
	m := make(map[int]bool, 0)
	for sum > 1 {
		sum = 0
		for i:=0;i<len(s);i++ {
			n,err := strconv.Atoi(string(s[i]))
			if err != nil {
				panic(err)
			}
			sum += n*n
		}
		if _,ok := m[sum];ok {
			sum = 0
		} else {
			m[sum] = false
		}
		s = strconv.Itoa(sum)
	}
	return sum
}