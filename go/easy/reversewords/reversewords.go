package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s\n", ReverseLine(scanner.Text()))
	}
}

func ReverseLine(in string) (out string) {
	out = ""
	inArray := strings.Split(in, " ")
	for i := len(inArray)-1; i >= 0; i-- {
		out += inArray[i] + " "
	}
	return
}
