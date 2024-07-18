/*
	Like libraries/modules in other languages, packages in Go support:
		modularity, encapsulation, separate compilation, and reuse.

	Usually the files of the xxx_xxx pacakge are stored in directory:
		$GOPATH/src/xxx_xxx
	( in my case it's /usr/local/go/src/xxx_xxx )

	By convention, a package's name matches the last segment of its【import path】
*/
package main 

import (
	"fmt"
	"xtina_ac"
)

func main() {
	// Each package serves as a separate [name space] for its declarations.
	// To refer to a function from outside its package, we must [qualify] the identifier:

	// the constant is declared in【2_packages_referee1.go】
	fmt.Println("---3---", xtina_ac.ExportedConst) 
	
	// fmt.Println(xtina_ac.notExportedConst) // compile error: undefined

	// the function is declared in【2_packages_referee2.go】
	xtina_ac.PrintExportedConst();
}

/*
	Any file may contain ANY number of【init】functions,
	（they're normal functions except that they can't be called or referenced）

	These【init】functions are automatically executed in the order 
	 in which they're declared when the program starts.
*/
func init() {
	fmt.Println("---1---")
}
func init() {
	fmt.Println("---2---")
}