package main

import (

	"os"
	"errors"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("Include file name as first parameter"))
	}
	filename := os.Args[1]

	fileinfo, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileinfo.Size())
}