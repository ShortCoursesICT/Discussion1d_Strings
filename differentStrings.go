package main

import (
	"fmt"
	"math"
)

func main() {
	stringA := `this is a string\n`
	stringA = `this is a new string\n`

	stringB := "this is a string\n"

	stringC := `\t\.xyz\n`
	stringD := "\\t\\.xyz\\n"

	fmt.Println(stringA)
	fmt.Println(stringB)
	fmt.Println(stringC)
	fmt.Println(stringD)
	math.Abs(10)
	//Good read on https://yourbasic.org/golang/gotcha-strings-are-immutable/

	/*	s := "hello"
		s[0] = 'H'
		fmt.Println(s)*/
	buf := []rune("hello") // converts to rune slice
	buf[0] = 'H'
	s := string(buf)
	fmt.Println(s) // "Hello"
}
