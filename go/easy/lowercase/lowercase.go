package main

import (

	"os"
	"errors"
	"bufio"
	"strings"
	"fmt"
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
		fmt.Println(strings.ToLower(line))
	}
}
