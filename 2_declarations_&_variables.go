/*
	There're 4 major kinds of declarations:
		1) var
		2) const
		3) type
		4) func
*/

package main

import (
	"fmt"
)

//【main】is a package-level declaration:
// 1) it is visible throughout all the files within the same package;
// 2) it is initialized BEFORE main begins. 
// (in contrast to local variables which are initialized when encountered during function execution)
func main() {
	/*
		the general form of a【var】declaration:
			【var name_of_var type_of_var = expression】

		either "type_of_var" or "= expression" may be omitted, but NOT BOTH;
		if the expression is omitted, the initial value is the ZERO value for the type
		(and thus in Go there is NO such things as as uninitialized variable):

		1) 0 for numbers
		2) false for booleans
		3) "" for strings
		4) nil for interfaces/reference types (slice/pointer/map/channel/function)
		5) for aggregate types (array/struct): ZERO values of all elements/fields
	*/
	var var1 string
	var var2 = ""
	fmt.Printf("%v, %v\n", var1, var2)

	// initialize a set of variables in a single declaration, with a matching list of expressions:
	// omitting the type allows declaration of multiple variables of different types
	var a, b, c bool = true, true, false
	var o, p, q = "ok", 1, true
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", a, b, c, o, p, q)
	
	//【short variable declarations】in the form of【name_of_var := expression】
	// (implicit type determination by the type of expression)
	i, j := 0, 0
	
	// they act like an【assignment】to variables already declared in the same lexical block
	j, k := 1, 2
	
	// (but at least one new variable must be declared therein)
	// i, j := 0, 1 // compile error: no new variables
	fmt.Printf("%v, %v, %v\n", i, j, k)
}