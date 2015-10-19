package main

import (
	"os"
	"errors"
	"bufio"
	"strconv"
	"fmt"
"math"
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
		p := Armstrong(line)
		l,err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if l == p {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}

func Armstrong(num string) int {
	sum := 0
	leng := len(num)
	for i:=0;i<len(num);i++ {
		l,err := strconv.Atoi(string(num[i]))
		if err != nil {
			panic(err)
		}
		sum += Pow(l, leng)
	}
	return sum
}

func Pow(base int, p int) int {
	return int(math.Pow(float64(base), float64(p)))
}