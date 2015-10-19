package main

import (
	"os"
	"errors"
	"bufio"
"strings"
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
		parts := strings.Split(line, ",")
		n,err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		m,err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		r := NmodM(n,m)
		fmt.Println(r)
	}
}

func NmodM(n int, m int) int {
	if n >= 0 && n < m {
		return n
	} else {
		return NmodM(n-m, m)
	}
}