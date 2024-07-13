/*
	Go has 25 reserved keywords that can't be used as names:
		break		default			func		interface		select
		case		defer			go 			map 			struct
		chan		else 			goto 		package 		switch
		const 		fallthrough 	if 			range 			type
		continue	for				import		return 			var
*/

package main

import (
	"fmt"
)

func main() {
	/*
		【tuple assignment】allows several variables to be assigned at once
		(all right-hand side expressions are evaluated before updating variables on the left-hand side)
		This will be very useful is some variables appear on both sides:
	*/

	// 1) swapping the values of 2 variables
	x := 1
	y := -1
	swap(&x, &y)
	fmt.Println(x, y)	// -1 1

	// 2) computing the GCD (greatest common divisor) of 2 integers
	fmt.Println(gcd_lcm(6, 21))	// 3 42

	// 3) assign unwanted values to the blank identifier
	gcd, _ := gcd_lcm(6, 21)
	fmt.Println(gcd)	// 3
}

func swap(p, q *int) {
	*p, *q = *q, *p
}

func gcd_lcm(x, y int) (int, int) {
	oldx, oldy := x, y
	for y != 0 {
		x, y = y, x%y
	}
	return x, (oldx * oldy) / x 
}