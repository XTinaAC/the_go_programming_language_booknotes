/*
	Use the【go doc】tool to access documents of the standard library from the command line
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// set the random seed (int64 nanoseconds since 1970)
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	fmt.Printf("random seed: %v \n", seed)

	// return a non-negative pseudo-random int in the half-open interval [0,n)
	tmp_int := rand.Intn(3)
	fmt.Printf("%s \n", coinflip1(tmp_int))
	
	heads := 0
	tails := 0
	// cases DO NOT fall through as in C
	switch coinflip2() {
		case "heads":
			heads++
		case "tails":
			tails++
		default:
			fmt.Println("landed on edge!")
	}
	fmt.Printf("heads-%d, tails-%d \n", heads, tails)
}

func coinflip1(x int) string {
	// a switch does not need an operand (a【tagless switch】; equivalent to【switch true】)
	switch {
		case x > 1:
			return "top"
		/* the optional DEFAULT case matches if none of the other cases does,
			and it may be placed ANYWHERE
		*/
		default:
			return "middle"
		case x < 1:
			return "bottom"
	}
}

func coinflip2() string {
	// return a pseudo-random float32 in the half-open interval [0.0,1.0)
	tmp_float32 := rand.Float32()
	fmt.Printf("random float32: %f \n", tmp_float32)
	if tmp_float32 < 0.5 {
		return "heads"
	} else {
		return "tails"
	}
}

