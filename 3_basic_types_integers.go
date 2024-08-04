/*
	Computers operate fundamentally on fixed-size numbers called words, which are interpreted as:
		1) integers
		2) floating-point numbers
		3) bit sets
		4) memory addresses

	Go's types fall into 4 categories:
		1) basic types（numbers, strings, booleans）
		2) aggregate types（arrays, structs）
		3) reference types（pointers, slices, maps, functions, channels）
		4) interface types

	(1) signed integers (numeric quantity):
		int8, int16, int32(rune; Unicode code point), int64
	(2) unsigned integers (raw data):
		uint8(byte), uint16, uint32, uint64

	-> Signed numbers are represented in 2's-complement form. 
		The range of values of an n-bit signed number is [ -2^(n-1), 2^(n-1) - 1 ].

	-> Unsigned integers have the range of [ 0, 2^n - 1 ].
*/

package main 

import (
	"fmt"
	"math"
)

func main() {
	/*
		Go's binary operators for arithmetic, logic, and comparison are listed below,
		 in order of decreasing precedence (top to bottom):

		 	*	/	%	<<	>>	&	&^
		 	+	-	|	^
		 	==	!=	<	<=	>	>=
		 	&&
		 	||

		Operators at the same level (on the same line) associate to the left, and thus parentheses () may be required for clarity.
		Each operator in the first 2 lines has a corresponding【assignment operator】(like  +=  &^=  >>=)
	*/

	/*
		Binary operators for arithmetic and logic (except shifts) MUST have operands of the SAME type.
		A conversion that narrows a BIG integer into a SMALLER one, or
			a conversion from interger to floating-point or vice versa,
			may change the value, or lose precision:
	*/	
	f := 3.14
	i := int(f)
	fmt.Println(f, i)	// 3.14	3

	/*
		【&】bitwise AND
		【|】bitwise OR
		【^】bitwise XOR
		【&^】bit clear

		For【z = x &^ y】，each bit of【z】equals:
			1)【0】if the corresponding bit of【y】is【1】;
			2)【x】otherwise.
	*/
	fmt.Println(0b0101 &^ 0b0011)	// 0b0100

	// The behavior of【/】depends on whether its operands are integers.
	fmt.Println(5/4.0, 5.0/4, 5/4)	// 1.25 1.25 1

	// The【remainder】operator【%】applies only to【integers】.
	// In Go, the sign of the【remainder】is the same as the sign of the【dividend被除数】.
	fmt.Println(-5%4, -5%-4)	// -1 -1

	/* 
		Integer literals of any size and type can be written as:
			1) ordinary decimal numbers
			2) [0b] binary numbers
			3) [0] octal numbers
			4) [0x or 0X] hexadecimal numbers
		We can control the radix and format with the [%d] [%b] [%o] and [%x] verbs.
	*/
	m := 0110
	fmt.Printf("%d %#[1]b %#[1]o %#[1]x \n", m) // 72 0b1001000 0110 0x48
	// The [1] adverb after % instructs Printf to use the 1st operand over and over again.
	// The # adverb for %o / %x / %X tells Printf to emit a 0 / 0x / 0X prefix respectively.

	// Print the range of values in hexadecimal:
	fmt.Printf("%x %x %x %x %x \n", math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64, math.MaxInt)
	fmt.Printf("%x %x %x %x %x \n", math.MinInt8, math.MinInt16, math.MinInt32, math.MinInt64, math.MinInt)
}