/*
	Find duplicate lines - version 1 - standard input
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create an empty map（of which【keys】are【strings】, and【values】are【ints】）
	counts := make(map[string]int)
	// create a variable that refers to a【bufio.Scanner】
	input := bufio.NewScanner(os.Stdin)
	// read the next line & remove the newline character, by calling【input.Scan()】
	for input.Scan() {
		// retrieve te result by calling【input.Text()】
		tmp := input.Text()
		if tmp == "end" { break }
		counts[tmp]++
	}
	for mapKey, count := range counts {
		if count > 1 {
			fmt.Printf("Found duplicates：【%s】appears【%d】times;\n", mapKey, count)
		}
	}
}