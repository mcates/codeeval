package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
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
		if strings.Contains(line, "|") {
			s := strings.Split(line, "|")
			enc := s[0]
			key := s[1]
			a := findWriter(enc, key)
			fmt.Println(a)
		}
	}
}

func findWriter(enc, key string) string {
	k := strings.Split(strings.TrimSpace(key), " ")
	var buffer bytes.Buffer
	for i := 0; i < len(k); i++ {
		i, err := strconv.Atoi(k[i])
		if err != nil {
			panic(err)
		}
		buffer.WriteString(string(enc[i-1]))
	}
	return buffer.String()
}
