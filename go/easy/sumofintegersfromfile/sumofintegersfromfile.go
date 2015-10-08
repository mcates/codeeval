package main

import (

	"os"
	"errors"
	"bufio"
	"strconv"
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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		val,err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum += val
	}
	fmt.Println(sum)
}
