package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val := NumOutcomes(line)
		fmt.Println(val)
	}
}

func NumOutcomes(line string) int {
	if len(line) != 0 && len(line) >= 2 && len(line) <= 70 && len(line)%4 == 0 {
		count := IsValid(line)
		return count
	} else {
		return 0
	}
}

func IsValid(vals string) int {
	i := 0
	mult := 1
	for i < len(vals) && mult > 0 {
		array := MakeArray(vals[i : i+4])
		toMult := Check(array)
		mult *= toMult
		i += 4
	}
	return mult
}

func Check(array []string) int {
	count := 0
	for _, value := range array {
		if value[0:2] == value[2:4] {
			count++
		}
	}
	return count
}

func MakeArray(part string) []string {
	i := 0
	vals := make([]string, 0)
	for i < len(part) {
		if string(part[i]) == "*" {
			if i == 0 {
				vals = append(vals, "A")
				vals = append(vals, "B")
			} else {
				typeA := "A"
				typeB := "B"
				length := len(vals)
				vals = append(vals, vals...)
				for index, value := range vals {
					if index < length {
						vals[index] = value + typeA
					} else {
						vals[index] = value + typeB
					}
				}
			}
		} else if string(part[i]) == "A" || string(part[i]) == "B" {
			if i == 0 {
				vals = append(vals, string(part[i]))
			} else {
				for j, value := range vals {
					vals[j] = value + string(part[i])
				}
			}
		} else {
			return []string{}
		}
		i++
	}
	return vals
}
