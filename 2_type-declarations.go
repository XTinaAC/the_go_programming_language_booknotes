/*
	The【type】of a variable / expression defines the characteristics of the values it may take on, such as:
	1) their [size] (number of bits / number of elements)；
	2) their [internal representation]；
	3) [intrinsic operations] that can be performed on them；
	4) [methods] associated with them。
*/

package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func main() {
	/*
		A【type declaration】defines a new【named type】that has the same【underlying type】as an existing type.
		It provides a way to separate different / incompatible uses of the【underlying type】so that they cannot be mixed unintentionally.
			type [name_of_type] [underlying_type]
	*/
	const (
		AbsoluteZeroC Celsius = -273.15
		FreezingC Celsius = 0
		BoilingC Celsius = 100
	)
	BoilingF, FreezingF := c_to_f(BoilingC), c_to_f(FreezingC)

	fmt.Printf("%g\n", BoilingC - FreezingC) // 100
	fmt.Printf("%g\n", BoilingF - FreezingF) // 180
	// fmt.Printf("%g\n", BoilingC - FreezingF) // compile error: invalid operation (mismatched types)

	// 基于相同底层类型的转换 (the [type] is changed, but not the [representation] of the value)
	// ( both variables have the same [underlying type],
	// 	or both are [unnamed pointer types] pointing to variables of the same [underlying type])
	fmt.Printf("%g\n", Celsius(BoilingF)) // 212

	// 基于不同底层类型的转换 (the [type] is changed, and so may the [representation] of the value)
	fmt.Printf("%d\n", int(BoilingF)) // 212

	/*
		some of the【conversion verbs】used by【Printf】:
			%d 			decimal integer
			%x %o %b 	integer in hexadecimal / octal / binary
			%f %g %e 	floating-point number
			%t 			boolean: true / false
			%c 			rune (Unicode code point)
			%s 			string
			%q 			quoted string "abc" or rune 'a'
			%v 			any value in a natural format
			%T 			type of any value
			%% 			literal percent sign (no operand)
	*/
}

func c_to_f(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

func f_to_c(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}