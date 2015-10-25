package main

import (
	"os"
	"errors"
	"bufio"
	"fmt"
"strings"
	"strconv"
	"sort"
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
		sorted := SimpleSort(line)
		for i:=0;i<len(sorted);i++ {
			if i==len(sorted)-1 {
				fmt.Printf("%.3f\n",sorted[i])
			} else {
				fmt.Printf("%.3f ", sorted[i])
			}
		}
	}
}

func SimpleSort(input string) []float64 {
	parts := strings.Split(input, " ")
	val := make([]float64, 0)
	for i:=0;i<len(parts);i++ {
		v,err := strconv.ParseFloat(parts[i], 64)
		if err != nil {
			panic(err)
		}
		val = append(val, v)
	}
	sort.Float64s(val)
	return val
}