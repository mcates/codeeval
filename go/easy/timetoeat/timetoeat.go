package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"sort"
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
		r := timeToEat(line)
		fmt.Println(r)
	}
}

func timeToEat(line string) string {
	parts := strings.Split(line, " ")
	sort.Strings(parts)
	var buffer bytes.Buffer
	for i := len(parts) - 1; i >= 0; i-- {
		buffer.WriteString(parts[i])
		if i != 0 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}
