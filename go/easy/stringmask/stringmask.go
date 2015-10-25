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
		r := stringMask(line)
		fmt.Println(r)
	}
}

func stringMask(line string) string {
	parts := strings.Split(line, " ")
	head := parts[0]
	mask := parts[1]
	var buf bytes.Buffer
	for i, v := range head {
		if b, _ := strconv.ParseBool(string(mask[i])); b {
			buf.WriteString(strings.ToUpper(string(v)))
		} else {
			buf.WriteString(string(v))
		}
	}
	return buf.String()
}

//func stringMask(line string) string { // using buffer resulted in lower memory usage
//	parts := strings.Split(line, " ")
//	head := parts[0]
//	mask := parts[1]
//	r := ""
//	for i, v := range head {
//		if b, _ := strconv.ParseBool(string(mask[i])); b {
//			r += strings.ToUpper(string(v))
//		} else {
//			r += string(v)
//		}
//	}
//	return r
//}
