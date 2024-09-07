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

	/*
		Strings are immutable sequences of bytes.
		A string shares the SAME underlying byte array with its【substrings】.
		Different【copies】of a string also share the SAME underlying memory.

		Unicode code point <=> rune <=> each assigned to a character in all of the world's writing systems,
	 	  plus accents and other diacritical makrs, control codes likes tab and carriage return, and plenty of esoterica.

		Unicode version 8 defines over 120,000 code points, and the ones that are in widespread use
		are fewer than 65,536 (which would fit in 16 bits)
		Go uses [ int32 ] to hold a single rune.

		UTF-8 is a [ variable-length encoding ] of Unicode code points as bytes, 
		and is now a Unicode standard. It uses between 1 & 4 bytes to represent each rune.
		The higher-order bits of the 1st byte of the encoding for a rune, 
		indicate how many bytes follow (which begin with 10):
			0xxxxxxx							0 ~ 127 (ASCII)
			110xxxxx 10xxxxxx					128 ~ 2047 (values < 128 are unused)
			1110xxxx 10xxxxxx 10xxxxxx			2048 ~ 65535 (values < 2048 are unused)
			11110xxx 10xxxxxx 10xxxxxx 10xxxxxx	65536 ~ 0x10ffff (other values unused)
		A rune whose value is less than 256 can be written with a single hexadecimal escape, such as:
			\x41	<=> 	A
		But for higher values, a \u (16-bit) or \U (32-bit) escape must be used:
	*/
	u0 := "A 世 界"
	u1 := "\x41 \xe4\xb8\x96 \xe7\x95\x8c"
	u2 := "\x41 \u4e16 \u754c"
	u3 := "\x41 \U00004e16 \U0000754c"
	fmt.Println(u0, u1, u2, u3)
}