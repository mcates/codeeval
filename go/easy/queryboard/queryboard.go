package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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

	board := make([]int, 256*256)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		Query(board, line)
	}
}

func Query(board []int, input string) {
	parts := strings.Split(input, " ")
	switch parts[0] {
	case "SetCol":
		j, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		x, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		for ii := 0; ii < 256; ii++ {
			board[j+256*ii-1] = x
		}
		break
	case "SetRow":
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		x, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		for ii := 0; ii < 256; ii++ {
			board[256*(i-1)+ii] = x
		}
		break
	case "QueryCol":
		sum := 0
		j, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for ii := 0; ii < 256; ii++ {
			sum += board[j+256*ii-1]
		}
		fmt.Println(sum)
		break
	case "QueryRow":
		sum := 0
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for ii := 0; ii < 256; ii++ {
			sum += board[256*(i-1)+ii]
		}
		fmt.Println(sum)
		break
	}
}
