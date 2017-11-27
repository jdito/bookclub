// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"strings"
)

func main() {}

func echo1(args []string) {
	// echo1 version
	s, sep := "", ""
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string) {
	// echo2 version
	s, sep := "", ""
	for _, arg := range args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	// echo3 version
	fmt.Println(strings.Join(args[0:], " "))
}

//!-
