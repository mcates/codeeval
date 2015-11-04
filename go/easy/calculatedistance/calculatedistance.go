package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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
		f := strings.SplitAfter(line, ")")
		p1 := strings.TrimRight(strings.TrimLeft(strings.TrimSpace(f[0]), "("), ")")
		p2 := strings.TrimRight(strings.TrimLeft(strings.TrimSpace(f[1]), "("), ")")
		s1 := strings.Split(p1, ",")
		s2 := strings.Split(p2, ",")
		i := calculateDistance(s1, s2)
		fmt.Println(i)
	}
}

func calculateDistance(p1, p2 []string) int {
	x1, err := strconv.Atoi(strings.TrimSpace(p1[0]))
	if err != nil {
		panic(err)
	}
	y1, err := strconv.Atoi(strings.TrimSpace(p1[1]))
	if err != nil {
		panic(err)
	}
	x2, err := strconv.Atoi(strings.TrimSpace(p2[0]))
	if err != nil {
		panic(err)
	}
	y2, err := strconv.Atoi(strings.TrimSpace(p2[1]))
	if err != nil {
		panic(err)
	}
	return int(math.Sqrt(float64((x1-x2)*(x1-x2) + ((y1 - y2) * (y1 - y2)))))
}
