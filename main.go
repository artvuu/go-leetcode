package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	name = "Gopher"
)

var programs = []any{
	program1,
}

func main() {
	argsWithProg := os.Args
	n, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	programs[n].(func())()

	fmt.Println(argsWithProg[1])
}

func program1() {
	fmt.Println("Hello World")
}
