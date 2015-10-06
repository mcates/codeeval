package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	firstline := true
	n := 0
	maplines := make(map[int]string)
	keys := make([]int,0) // check that value doesn't already exist
	for scanner.Scan() {
		if firstline {
			n, _ = strconv.Atoi(scanner.Text())
			firstline = false
		} else {
			maplines[len(scanner.Text())] = scanner.Text()
			keys = append(keys, len(scanner.Text()))
		}
	}
	sort.Ints(keys)
	for max := len(keys); max > len(keys)-n; max-- {
		fmt.Println(maplines[keys[max-1]])
	}
}
