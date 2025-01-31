/*
   【复合数据类型】
   	[ Basic types ] are to [ composite types ], 
   		like the atoms to the molecules of our universe.
   		
   	(1) Arrays (homogeneous) & (2) structs (heterogeneous) are fixed size;
   	(3) Slices & (4) maps are dynamic structures that can grow and shrink.
*/

package main 

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	/*
		Arrays are rarely used directly in Go due to their【fixed length】.
	*/
	/*
		Elements of a new array variable are initially set to the【ZERO value】
			for the element type, which is 0 for numbers.
	*/
	var a1 [3]int = [3]int {111, 222}
	for idx, val := range a1 {
		fmt.Println(idx, val)	//	0 111		1 222		2 0
	}
	/*
		When initializing an array with an【array literal】, if an ellipsis...
			appears in place of the【size】, the array length is determined by the number of initializers.
	*/
	a2 := [...]int {111, 222, 333, 444}
	fmt.Printf("%T\n", a2)	// [4]int

	/*
		The【size】is part of the array type, and it must be a【constant expression】.
		It is also possible to specify a list of【index】&【value】pairs.
	*/
	type Currency int
	const (
		RMB Currency = iota	// constant generator
		USD
		EUR
		GBP
	)
	a3 := [1 + GBP]string {RMB: "¥", USD: "$", EUR: "€", GBP: "£"}
	fmt.Printf("%T\n", a3)	// [4]string
	fmt.Println(GBP, a3[GBP])	// 3 £

	/*
		In this formm indices can appear in ANY order,
			and some may be omitted:
	*/
	a4 := [...]int {50:-1, 1:-1, 99: -1}
	fmt.Println(a4)
	// an array with 100 elements
	// [ 0 -1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//   0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 -1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//   0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 -1 ]


	/*
		The array type is comparable if its element type is comparable.
	*/
	b1 := [2]int {1, 2}
	b2 := [...]int {1, 2}
	fmt.Println(b1==b2, b1!=b2)	// true false

	/*
		As a more plausible example, the function Sum256/Sum384/Sum512 (32/48/64 bytes respectively)
		produces the cryptographic hash/digest of a message stored in an arbitrary byte slice.
	*/
	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("A"))
	fmt.Printf(" %x \n %x \n %t \n %T \n", c1, c2, c1==c2, c1)
	// ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb 
	// 559aead08264d5795d3909718cdd05abd49572e84fe55590eef31a88a08fdffd 
	// false 
	// [32]uint8

	// ------ Passing arrays as arguments ------
	fmt.Printf("before::: %x \n", c1)
	zero(c1)
	fmt.Printf("after::: %x \n", c1)
	// before::: ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb 
	// after::: ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb 

	fmt.Printf("before::: %x \n", c1)
	zero_1(&c1)
	fmt.Printf("after::: %x \n", c1)
	// before::: ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb 
	// after::: 0000000000000000000000000000000000000000000000000000000000000000

	fmt.Printf("before::: %x \n", c2)
	zero_2(&c2)
	fmt.Printf("after::: %x \n", c2)
	// before::: 559aead08264d5795d3909718cdd05abd49572e84fe55590eef31a88a08fdffd 
	// after::: 0000000000000000000000000000000000000000000000000000000000000000 
}

/*
	Different from languages that implicitly pass arrays [by reference],
	Go treats arrays like any other type: functions receive a copy of each
	argument, not the original.
*/
func zero(arr [32]byte) [32]byte {
	for i := range arr {
		arr[i] = 0
	}
	return arr 	// 0000000000000000000000000000000000000000000000000000000000000000
}
/*
	We can explicitly pass a pointer to an array so that any modifications
	the function makes to array elements will be visible to the caller.
*/
func zero_1(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0 // set each byte to 0
	}
}
func zero_2(ptr *[32]byte) {
	*ptr = [32]byte {}
}