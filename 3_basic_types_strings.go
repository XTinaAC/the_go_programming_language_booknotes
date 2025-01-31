/*
	Text strings are conventionally interpreted as 
		UTF-8-encoded sequences of 
		Unicode code points (runes).
*/

package main 

import (
	"fmt"
	"unicode/utf8"
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

		Unicode code point <=> [ rune ] <=> each assigned to a character in all of the world's writing systems,
	 	  plus accents and other diacritical makrs, control codes likes tab and carriage return, and plenty of esoterica.

		Unicode version 8 defines over 120,000 code points, and the ones that are in widespread use
		are fewer than 65,536 (which would fit in 16 bits)
		Go uses [ int32 ] to hold a single [ rune ].

		UTF-8 is a [ variable-length encoding ] of Unicode code points as bytes, 
		and is now a Unicode standard. It uses between 1 & 4 bytes to represent each rune.

		The higher-order bits of the 1st byte of the encoding for a rune, 
		indicate how many bytes follow (which begin with 10):
			0xxxxxxx				0 ~ 127 (ASCII)
			110xxxxx 10xxxxxx			128 ~ 2047 (values < 128 are unused)
			1110xxxx 10xxxxxx 10xxxxxx		2048 ~ 65535 (values < 2048 are unused)
			11110xxx 10xxxxxx 10xxxxxx 10xxxxxx	65536 ~ 0x10ffff (other values unused)

		A rune whose value is less than 256 can be written with a single hexadecimal escape, such as:
			\x41	<=> 	A
		But for higher values, a【\u (16-bit)】or【\U (32-bit)】escape must be used.
		（下面示例：Unicode escapes that denotes the UTF-8 encoding of the specified numeric code point value）
	*/

	// "A" 对应的numeric code point【16进制】为【\u41】(或【\U00000041】)，对应【2进制】为【0100 0001】；
	// 		1-byte unicode编码为【0100 0001】，对应【16进制】为【\x41】

	// "世" 对应的numeric code point【16进制】为【\u4e16】(或【\U00004e16】)，对应【2进制】为【0100 1110 0001 0110】；
	// 		3-byte unicode编码为【1110 0100 1011 1000 1001 0110】，对应【16进制】为【\xe4 \xb8 \x96】

	// "界" 对应的numeric code point【16进制】为【\u754c】(或【\U0000754c】)，对应【2进制】为【0111 0101 0100 1100】；
	// 		3-byte unicode编码为【1110 0111 1001 0101 1000 1100】，对应【16进制】为【\xe7 \x95 \x8c】

	u0 := "A 世 界" 
	u1 := "\x41 \xe4\xb8\x96 \xe7\x95\x8c"	// 1110 
	u2 := "\u0041 \u4e16 \u754c"	// x0100 
	u3 := "\U00000041 \U00004e16 \U0000754c"
	fmt.Println("\n", u0, "\n", u1, "\n", u2, "\n", u3)

	fmt.Println(len(u0))	// number of bytes: 9
	fmt.Println(utf8.RuneCountInString(u0))	// number of runes (code points): 5

	// explicit decoding
	for i := 0; i < len(u0); {
		r, size := utf8.DecodeRuneInString(u0[i:])
		fmt.Printf("%d \t %c \n", i, r)
		i += size
	}

	// implicit decoding
	for i, r := range u0 {
		fmt.Printf("%d \t %q \t %d \n", i, r, r)
	}
	// 打印结果如下：
	// 0 	 'A' 	 65 
	// 1 	 ' ' 	 32 
	// 2 	 '世' 	 19990 
	// 5 	 ' ' 	 32 
	// 6 	 '界' 	 30028

	// Each time a UTF-8 decoder, whether explicit in a call to utf8.DecodeRuneInString or implicit in a range loop,
	// consumes an unexpected input byte, it generates a special Unicode replacement character, '\uFFFD', 
	// which is usually printed as a white question mark inside a black hexagonal or diamond-like shape �.
	// When a program encounters this rune value, it’s often a sign that some upstream part of the system 
	// that generated the string data has been careless in its treatment of text encodings.
	fmt.Println("\ufffd")

	// %d 	decimal integer
	// %x, %o, %b 	integer in hexadecimal, octal, binary
	// %f, %g, %e 	floating-point number: 3.141593 3.141592653589793 3.141593e+00
	// %t 	boolean: true or false
	// %c 	rune (Unicode code point)
	// %s 	string
	// %q 	quoted string "abc" or rune 'c'
	// %v 	any value in a natural format
	// %T 	type of any value
	// %% 	literal percent sign (no operand)

	// //【中文范围】\u4e00 (十进制：19968) - \u9fa5 (十进制：40869) 
	// for i := 19968; i <= 40869; {
	// 	fmt.Printf("%c ", i)
	// 	i += 1
	// }
}