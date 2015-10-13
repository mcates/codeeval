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
		parts := strings.Split(line, ",")
		head := parts[0]
		tail := parts[1]
		index := Position(head, tail)
		fmt.Println(index)
	}
}

func Position(h string, t string) int {
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] == t[0] {
			return i
		}
	}
	return -1
}
