package main

import (
	"os"
	"errors"
	"bufio"
"strings"
	"fmt"
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
		r := Beautiful(line)
		fmt.Println(r)
	}
}

func Beautiful(input string) int {
	m := make(map[string]int, 0)
	for i := 0; i < len(input); i++ {
		v := strings.ToLower(string(input[i]))
		if v >= "a" && v <= "z" {
			if _,ok := m[v];ok {
				m[v] = m[v] + 1
			} else {
				m[v] = 1
			}
		}
	}
	sorted := make([]int, 0)
	for _,value := range m {
		sorted = append(sorted, value)
	}
	sort.Ints(sorted)
//	fmt.Printf("%+v\n", sorted)
	sum := 0
	for i := len(sorted) - 1; i >=0; i-- {
		sum += (26 - (len(sorted) - 1) + i) * sorted[i]
	}
	return sum
}