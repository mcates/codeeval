package main

import (
	"bufio"
	"bytes"
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
		r := swapCase(line)
		fmt.Println(r)
	}
}

func swapCase(line string) string {
	var buffer bytes.Buffer
	for _, v := range line {
		switch {
		case "a" <= string(v) && string(v) <= "z":
			buffer.WriteString(strings.ToUpper(string(v)))
			break
		case "A" <= string(v) && string(v) <= "Z":
			buffer.WriteString(strings.ToLower(string(v)))
		default:
			buffer.WriteString(string(v))
			break
		}
		//		fmt.Printf("%v, %v ", i, string(v))
	}
	return buffer.String()
}
