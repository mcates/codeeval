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
	for scanner.Scan() {
		line := scanner.Text()
		sum := SumOfDigits(line)
		fmt.Println(sum)
	}
}

func SumOfDigits(line string) int {
	sum := 0
	for i:=0;i<len(line);i++ {
		val,_ := strconv.Atoi(string(line[i]))
		sum += val
	}
	return sum
}