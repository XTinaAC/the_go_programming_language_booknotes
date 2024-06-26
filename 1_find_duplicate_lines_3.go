/*
	Find duplicate lines - version 3 - a list of file names
		"one-gulp" mode - read the entire file input into memory 
						  & split it into lines all at once, all before processing
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//【ioutil.ReadFile】returns（1）a byte slice and（2）the built-in error type
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
		// convert the byte slice into a string, and split it using【strings.Split】
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	// the order of map iteration is RANDOM（by intentional design）
	for mapKey, count := range counts {
		if count > 1 {
			fmt.Printf("Found duplicates:【%s】appears【%d】times;\n", mapKey, count)
		}
	}
}