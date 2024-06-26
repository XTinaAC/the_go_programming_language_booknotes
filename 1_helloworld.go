// By convention, the first line is commentary describing the package (or the program as a whole for a main package);

package main

// 2 ways to import multiple packages:

// 1) individual imports
// import "os"
// import "fmt"

// 2) a parenthesized list
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello, world!")

	// If not explicitly initialized, variables are implictly initialized to the ZERO value of its type
	// (0 for numeric types and the empty string "" for strings)
	var s, tmp string

	// "i++" is a statement (not an expression as in C) and is postfix only (so "j = i++" and "++i" are illegal)
	for i := 1; i < len(os.Args); i++ {
		// Each arithmetic & logical operator has a corresponding assignment operator (like += *=)
		s += tmp + os.Args[i]
		tmp = ", "
	}
	fmt.Println("You've typed:" + s)

	// ":=" is part of a short variable declaration 
	s1, tmp1 := "", ""

	// Another form of the for loop iterates over a range of values
	for idx, val := range os.Args[1:] {
		// String concatenation could be costly if the amount of data involved is large.
		s1 += tmp1 + val
		tmp1 = ", "
		fmt.Printf("序号%d、", idx) 
	}
	fmt.Println("\nYou've types:" + s1)

	// A simpler and more efficient solution:
	fmt.Println(strings.Join(os.Args[1:], ", "))
	
	// We can also print any slice this way:
	fmt.Println(os.Args[1:]) 
}
