/*
	Go provides 2 sizes of floating-point numbers:
		【float32】 & 【float64】
	arithmetic properties of which are governed by the IEEE 754 standard.
*/

package main 

import (
	"fmt"
	"math"
)

func main() {
	/*
		The approximate number of digits of precision:
			1) float64: 15 digits
			2) float32:  6 digits

		The smallest positive values are near:
			1) float64: 4.9e-324
			2) float32: 1.4e-45

		And the max values are as follows.	
	*/
	fmt.Printf("%g %g \n", math.MaxFloat32, math.MaxFloat64)	// about 3.4e38 1.8e308
	
	/*
		(【%g】 chooses the most compact representation that has adequate precision )
		(【%f】 (np exponent) )	
		(【%e】 (exponent) )
		( all 3 verbs allows [field width] and [numeric precision]: e.g., %8.3f )
	*/
	
	// The smallest positive integers that CANNOT be represented EXACTLY:
	var t1 float32 = 1 << 24;
	var t2 float64 = 1 << 53;
	fmt.Println(t1 == t1+1, t2 == t2+1)	// true true
	fmt.Println(t1 == t1-1, t2 == t2-1)	// false false

	/*
		Special values defined by IEEE 754:
			1)【 +Inf -Inf 】positive & negative infinities
				( numbers of excessive megnitude / result of division by 0 )
			2)【 NaN 】not a number
				( result of such mathematically dubious operations as 0/0 or Sqrt(-1) )
	*/
	var zero float64
	p_inf := math.Inf(0)
	n_inf := math.Inf(-1)
	nan := math.NaN()
	fmt.Println(p_inf, n_inf, 1/zero, -1/zero)	// +Inf -Inf +Inf -Inf
	fmt.Println(nan, 0/zero, math.Sqrt(-1))		// NaN NaN NaN

	// equal comparision with【NaN】、【Inf】:
	fmt.Println(p_inf == p_inf, n_inf == n_inf)	// true true
	fmt.Println(nan == nan)	// false; // any comparision with【NaN】yields false
}