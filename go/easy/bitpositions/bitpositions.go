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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := strings.Split(line, ",")
		first := parsed[0]
		asInt, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}
		n := int64(asInt)
		second, err := strconv.Atoi(parsed[1])
		if err != nil {
			panic(err)
		}
		third, err := strconv.Atoi(parsed[2])
		if err != nil {
			panic(err)
		}
		result := CompareBitPositions(strconv.FormatInt(n, 2), second, third)
		fmt.Println(result)
	}
}

func CompareBitPositions(main string, p1 int, p2 int) string {
	if main[len(main)-p1] == main[len(main)-p2] {
		return "true"
	} else {
		return "false"
	}
}
