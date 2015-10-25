package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
		r := CapitalizeWords(line)
		fmt.Println(r)
	}
}

func CapitalizeWords(line string) string {
	parts := strings.Split(line, " ")
	for i, v := range parts {
		f := string(v[0])
		f = strings.ToUpper(f)
		parts[i] = f + v[1:]
	}
	c := strings.Join(parts, " ")
	return c
}
