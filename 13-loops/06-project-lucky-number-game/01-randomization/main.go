// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Seed(10)
	// rand.Seed(100)

	// t := time.Now()
	// rand.Seed(t.UnixNano())

	// ^-- same:

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	guess := 10

	for n := 0; n != guess; {
		n = randGen.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
