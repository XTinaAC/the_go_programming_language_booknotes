/*
	Text strings are conventionally interpreted as 
		UTF-8-encoded sequences of 
		Unicode code points (runes).
*/

package main 

import (
	"fmt"
)

func main() {
	/*
		The built-in【len】function returns the number of【bytes】(NOT runes) in a string;
		The index operation【s[i]】retrieves the i-th byte of string s (0 <= i <= len(s));
		The UTF-8 encoding of a non-ASCII code point requires >=2【bytes】.
	*/
	s := "a和b"
	fmt.Println(len(s)) // 5
	fmt.Println(s[0], s[1], s[2], s[3], s[4]) // 97 229 146 140 98

	// The substring operation【s[i:j]】starts at the byte at index i, 
	// and continues up to (but not including) the byte at index j
	fmt.Println(s[1:4]) // 和

	// Strings are immutable sequences of bytes.
	// A string shares the SAME underlying byte array with its【substrings】.
	// Different【copies】of a string also share the SAME underlying memory.
}