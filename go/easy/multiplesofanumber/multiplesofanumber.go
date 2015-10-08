package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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
		parsed := strings.Split(line, ",")
		head := parsed[0]
		tail := parsed[1]
		smalletMultiple := SmallestMultiple(head, tail)
		fmt.Println(smalletMultiple)
	}
}

func SmallestMultiple(head string, tail string) int {
	first,err := strconv.Atoi(head)
	if err != nil {
		panic(err)
	}
	second,err := strconv.Atoi(tail)
	if err != nil {
		panic(err)
	}
	multiple := second
	for multiple < first {
		multiple += second
	}
	return multiple
}