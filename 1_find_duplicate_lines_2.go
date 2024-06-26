/*
	Find duplicate lines - version 2 - standard input / a list of file names
		"streaming" mode - input is read & broken into lines as needed
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range filenames {
			//【os.Open】returns（1）an open file（*os.File）and（2）the built-in error type
			file, err := os.Open(filename)

			if err != nil {
				// Print the error message onto the standard error stream using【Fprintf】
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				continue
			}

			countLines(file, counts)
			// close the file and release any related resources
			file.Close()
		}
	}
	for mapKey, count := range counts {
		if count > 1 {
			fmt.Printf("Found duplicates:【%s】appears【%d】times;\n", mapKey, count)
		}
	}
}

// Functions & other package-level entities can be declared in ANY order
func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		tmp := input.Text()
		if tmp == "end" { break }

		// any changes made in the callee's copy of the map reference, 
		// will be visible through the caller's map reference too
		counts[tmp]++
	}
}