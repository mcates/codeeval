package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"bytes"
)

type StringValid struct {
	Value string
	CanReplace bool
}

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		head := parts[0]
		tail := parts[1]
		pairs := strings.Split(tail, ",")
		i := 0
		toReplace := make([]string, 0)
		replaceWith := make([]string, 0)
		for i < len(pairs) {
			toReplace = append(toReplace, pairs[i])
			replaceWith = append(replaceWith, pairs[i+1])
			i = i + 2
		} // finish parsing strings
		output := Substitute(head, toReplace, replaceWith)
		fmt.Println(output)
	}
}

func Substitute(head string, toReplace []string, replaceWith []string)  string {
	replaced := make([]StringValid, 0) // replaced index, length
	canReplace := StringValid{head, true}
	replaced = append(replaced, canReplace)
	i := 0
	for i < len(toReplace) { // look through each replace string
		replaced = FindInstances(toReplace[i], replaceWith[i], replaced)
		i++
	}
	j := 0
	var buffer bytes.Buffer
	for j < len(replaced) {
		buffer.WriteString(replaced[j].Value)
		j++
	}
	return buffer.String()
}

func FindInstances(toReplace string, replaceWith string, inputMap []StringValid) []StringValid {
	i := 0
start:
	for i < len(inputMap) {
		if inputMap[i].CanReplace {
			if len(toReplace) <= len(inputMap[i].Value) {
				if strings.Contains(inputMap[i].Value, toReplace) {
					j := 0
					for j < len(inputMap[i].Value) {
						if toReplace == inputMap[i].Value[j:j+len(toReplace)] {
							if len(toReplace) == len(inputMap[i].Value) {
								inputMap[i] = StringValid{replaceWith, false}
								goto start
							} else
							// if at beginning of string
							if strings.HasPrefix(inputMap[i].Value, toReplace) {
								endInputMap := inputMap[i+1:]
								beginningInputMap := inputMap[:i]
								middleInputMap := make([]StringValid, 0)
								middleInputMap = append(middleInputMap, StringValid{replaceWith, false})
								middleInputMap = append(middleInputMap, StringValid{inputMap[i].Value[len(toReplace):], true})
								finalInputMap := append(middleInputMap, endInputMap...)
								inputMap = append(beginningInputMap, finalInputMap...)
								goto start
							} else
							// if at end of string
							if strings.HasSuffix(inputMap[i].Value, toReplace) {
								endInputMap := inputMap[i+1:]
								beginningInputMap := inputMap[:i]
								middleInputMap := make([]StringValid, 0)
								middleInputMap = append(middleInputMap, StringValid{inputMap[i].Value[:j], true})
								middleInputMap = append(middleInputMap, StringValid{replaceWith, false})
								finalInputMap := append(middleInputMap, endInputMap...)
								inputMap = append(beginningInputMap, finalInputMap...)
								goto start
							} else {
								// if in middle of string
								endInputMap := inputMap[i+1:]
								beginningInputMap := inputMap[:i]
								middleInputMap := make([]StringValid, 0)
								middleInputMap = append(middleInputMap, StringValid{inputMap[i].Value[:j], true})
								middleInputMap = append(middleInputMap, StringValid{replaceWith, false})
								middleInputMap = append(middleInputMap, StringValid{inputMap[i].Value[j+len(toReplace):], true})
								finalInputMap := append(middleInputMap, endInputMap...)
								inputMap = append(beginningInputMap, finalInputMap...)
								goto start
							}
						}
						j++
					}
				}
			}
		}
		i++
	}
	return inputMap
}

