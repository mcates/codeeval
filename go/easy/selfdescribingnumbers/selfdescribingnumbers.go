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
		i := SelfDescribing(line)
		fmt.Println(i)
	}
}

func SelfDescribing(s string) int {
	for i:=0;i<len(s);i++ {
		n := s[i]
		nInt,err := strconv.Atoi(string(n))
		if err != nil {
			panic(err)
		}
		total := 0
		for j:=0;j<len(s);j++ {
			if strconv.Itoa(i) == string(s[j]) {
				total++
			}
		}
		if total != nInt {
			return 0
		}
	}
	return 1
}