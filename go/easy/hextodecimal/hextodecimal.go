package main

import (
	"os"
	"errors"
	"bufio"
	"fmt"
	"strconv"
"math"
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
		d := Hex(line)
		fmt.Println(d)
	}
}

func Hex(line string) int {
	sum := 0
	for i:=len(line)-1;i>=0;i-- {
		switch string(line[i]) {
		case "a", "A":
			sum += 10 * Pow(16,len(line)-1-i)
			break
		case "b","B":
			sum += 11 * Pow(16,len(line)-1-i)
			break
		case "c","C":
			sum += 12 * Pow(16,len(line)-1-i)
			break
		case "d","D":
			sum += 13 * Pow(16,len(line)-1-i)
			break
		case "e","E":
			sum += 14 * Pow(16,len(line)-1-i)
			break
		case "f","F":
			sum += 15 * Pow(16,len(line)-1-i)
			break
		default:
			v,err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			sum += v *  Pow(16,len(line)-1-i)
			break
		}
	}
	return sum
}

func Pow(base int, p int) int {
	return int(math.Pow(float64(base), float64(p)))
}