package main

import (

	"os"
	"errors"
	"bufio"
"strings"
	"fmt"
	"strconv"
	"sort"
	"bytes"
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
		ret := UniqueElements(line)
		fmt.Println(ret)
	}
}

func UniqueElements(line string) string {
	parts := strings.Split(line,",")
	r := make(map[int]int, 0)
	for i:=0;i<len(parts);i++ {
		cur,_ := strconv.Atoi(parts[i])
		r[cur] = cur
	}
	a := make([]int, 0)
	for k,_ := range r {
		a = append(a, k)
	}
	sort.Ints(a)
	var buffer bytes.Buffer
	for j:=0;j<len(a);j++ {
		if j == len(a)-1 {
			buffer.WriteString(strconv.Itoa(a[j]))
		} else {
			buffer.WriteString(strconv.Itoa(a[j]))
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}