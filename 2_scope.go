/*
	The【scope】of a declaration is a region of the program text. 
		It's a compile-time property.
	The【lifetime】of a variable is the range of time during execution when it can be referred to by other parts of the program.
		It's a run-time property.

	There's a【lexical block】for:
	1) the entire source code (the【universal block】)
	2) each package
	3) each file
	4) each for / if / switch statement; and each case in a switch / select statement
	5) each explicit【syntactic block】(a sequence of statements enclosed in braces)

	e.g.:
	1) built-in types / functions / constants (like int, len, true) -> the【universal block】
	2) declarations outside any function -> at【package level】
	3) declarations of imported packages -> at【file level】
	4) 5) local declarations

	A program may contain multiple declarations of the [same name] so long as
		each declaration is in a different lexical block.
	Inner declarations will shadow/hide outer ones, making them inaccessible.
*/

package main 

import (
	"fmt"
)

func main() {
	/*
		The following example illustrates scope rules (not good style though):

		The【for】loop creates 2 lexical blocks:
		1) an explicit block for the loop body;
		2) an implicit block for the initialization clause.

		The example below has 3 variables named [idx1], each declared in a different block.
		Beware that the 2nd【short varible declartion】of [idx1] makes the 1st one inaccessible
			and thus the statement DOES NOT update the 1st [idx1].
	*/
	x := "hello"
	idx1 := 0
	for _, idx1 := range x {
		idx1 := idx1 + 'A' - 'a'
		fmt.Printf("%c", idx1)	// HELLO
	}
	fmt.Println("\n", idx1)	// 0

	// The most direct way to change this, is to
	// replace the 2nd【declaration】with【assignment】:
	idx2 := 0
	for ; idx2 < len(x); idx2++ {
		idx2 := x[idx2] + 'A' - 'a'
		fmt.Printf("%c", idx2)	// HELLO
	}
	fmt.Println("\n", idx2)	// 5

	/* 
		Like【for loops】,【if】statements and【switch】statements also create 
		implicit blocks in addition to their body blocks.

		The 2nd if statement is NESTED within the 1st, so variables declared within the 1st
		statement's initializer are visible within the 2nd.
	*/
	if y := 1; y == 0 {
		fmt.Println(y)
	} else if z := 2; z == y {
		fmt.Println(y, z)
	} else {
		fmt.Println(y, z)	// 1 2
	}

	/*
		Same rules apply to each case of a switch statement.
		There's an implicit block for the condition and an explicit one for each case body.
	*/
	q := 1
	switch q := q+1; q {
		case 1:
			m := q + 1
			fmt.Println(q, m)
		case 2:
			m := q + 1
			fmt.Println(q, m) // 2 3
	}
}