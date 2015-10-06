package main

import (
	"bufio"
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
		output, err := FizzBuzz(scanner.Text())
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	}
}

func FizzBuzz(input string) (string, error) {
	fizzbuzzreturn := ""
	err := ""
	parsed := strings.Fields(input)
	fizz, _ := strconv.Atoi(parsed[0])
	if fizz > 20 || fizz < 1 {
		err = "Fizz is out of bounds.  range [1, 20]\n"
	}
	buzz, _ := strconv.Atoi(parsed[1])
	if buzz > 20 || buzz < 1 {
		err += "Buzz is out of bounds.  range [1, 20]\n"
	}
	length, _ := strconv.Atoi(parsed[2])
	if length < 21 || length > 100 {
		err += "Length is out of bounds.  range [21, 100]\n"
	}
	if len(err) > 1 {
		return fizzbuzzreturn, errors.New(err)
	}
	for i := 1; i <= length; i++ {
		if i%fizz == 0 && i%buzz == 0 {
			fizzbuzzreturn += "FB "
		} else if i%fizz == 0 {
			fizzbuzzreturn += "F "
		} else if i%buzz == 0 {
			fizzbuzzreturn += "B "
		} else {
			fizzbuzzreturn += strconv.Itoa(i) + " "
		}
	}
	return fizzbuzzreturn, nil
}
