// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
//
// 3. After you've done, check out the solution
//    and read the comments there
//
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// ---------------------------------------------------------

func main() {
	var (
		i    int
		i8   int8
		i16  int16
		i32  int32
		i64  int64
		f32  float32
		f64  float64
		c64  complex64
		c128 complex128
		b    bool
		s    string
		r    rune
		B    byte
	)

	fmt.Println(i)
	fmt.Println(i8)
	fmt.Println(i16)
	fmt.Println(i32)
	fmt.Println(i64)
	fmt.Println(f32)
	fmt.Println(f64)
	fmt.Println(c64)
	fmt.Println(c128)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(r)
	fmt.Println(B)
}
