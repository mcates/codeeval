package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
	"bytes"
	"sort"
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
		arr := strings.Split(line, ";")
		head := arr[0]
		tail := arr[1]
		sh := strings.Split(head, ",")
		st := strings.Split(tail, ",")
		r := Intersect(sh, st)
		var buffer bytes.Buffer
		for i:=0;i<len(r);i++{
			if i == len(r)-1 {
				buffer.WriteString(strconv.Itoa(r[i]))
			} else {
				buffer.WriteString(strconv.Itoa(r[i]))
				buffer.WriteString(",")
			}
		}
		fmt.Println(buffer.String())
	}
}

func Intersect(h []string, t []string) []int {
	mh := make(map[string]bool, 0)
	for i:=0;i<len(h);i++ {
		mh[h[i]] = false
	}
	for j:=0;j<len(t);j++ {
		if _,ok := mh[t[j]]; ok {
			mh[t[j]] = true
		}
	}
	ia := make([]int, 0)
	for k,v := range mh {
		if v {
			asInt,err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			ia = append(ia, asInt)
		}
	}
	sort.Ints(ia)
	return ia
}