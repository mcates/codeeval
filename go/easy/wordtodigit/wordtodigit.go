package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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
		r := WordToDigit(line)
		fmt.Println(r)
	}
}

func WordToDigit(input string) string {
	r := ""
	parts := strings.Split(input, ";")
	for i := 0; i < len(parts); i++ {
		switch parts[i] {
		case "zero":
			r += "0"
			break
		case "one":
			r += "1"
			break
		case "two":
			r += "2"
			break
		case "three":
			r += "3"
			break
		case "four":
			r += "4"
			break
		case "five":
			r += "5"
			break
		case "six":
			r += "6"
			break
		case "seven":
			r += "7"
			break
		case "eight":
			r += "8"
			break
		case "nine":
			r += "9"
			break
		}
	}
	return r
}
