/*
	Every variable has an address. Variables are addressable values.
*/

package main

import (
	"fmt"
	"strings"
	"flag"
)

func main() {
	/*
		&x (address of x) yields【a pointer to x】which is of type【*int】
			and【*p】yields the value of x
	*/
	var x int
	var p = &x
	fmt.Println(*p) // 0

	/*
		 The【new(T)】expression:
		 1) creates an unnamed variable of type T
		 2) initializes it to the ZERO value of T
		 3) returns its address that is of type *T
	*/
	m := new(int)
	fmt.Println(*m)	// 0

	// Since【*m】denotes a variable, it may also
	// appear on the left-hand side of an assignment
	*m = 1
	fmt.Println(*m)	// 1

	// Two pointers are equal if & only if 
	// 1) they point to the same variable
	// 2) both are nil (the ZERO value for a pointer)
	a := newInt1()
	b := newInt1()
	fmt.Println(a==b) // false

	fmt.Println(delta(20, 80))	// 60

	/*
		Pointers are key to the【flag】package 
		(it uses a program's command-line arguments to set the values of variables within it)

		【flag.Bool】creates a flag variable of type【bool】, and returns a pointer to it
		【flag.String】creates a flag variable of type【string】, and returns a pointer to it

		They take 3 arguments:
		1) the [name] of the flag
		2) the flag variable's [default value]
		3) a [message] to be printed when users provide an [invalid argument] / [invalid flag] / [-h] or [-help]
	*/
	var omitNewLine = flag.Bool("onl", false, "Omit trailing newline")
	var separator = flag.String("spt", " ", "The separator string")

	/*
		【https://pkg.go.dev/flag#Parse】
		【flag.parse】parses the command-line flags from os.Args[1:], 
			and update the flags from their default values. It must be called:
		1) after all flags are defined;
		2) before any flag is accessed by the proram.
	*/
	flag.Parse()

	// Access values of flags INDIRECTLY as【*omitNewLine】and【*separator】 
	fmt.Print(strings.Join(flag.Args(), *separator))
	if !*omitNewLine {
		fmt.Println()
	}
	/* 
	  (eg.)
		go run 2_pointers\&new.go -h
		go run 2_pointers\&new.go A B C
		go run 2_pointers\&new.go -onl A B C
		go run 2_pointers\&new.go -spt "; " A B C
	*/
}

// The following 2 newInt functions have IDENTICAL behaviors
// For the【new()】case, there's no need to declare a dummy name
func newInt1() *int {
	return new(int)
}
func newInt2() *int {
	var dummy int
	return &dummy
}

//【new】is a predeclared function, not a keyword
// and thus we can redefine the name for sth. else 
// (of course the built-in new function is unavailable here)
func delta(old, new int) int { return new - old }