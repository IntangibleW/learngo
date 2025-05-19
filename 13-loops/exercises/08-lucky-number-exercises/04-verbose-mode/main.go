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
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

var verbose bool

func main() {
	if os.Args[1] == "-v" {
		verbose = true
	}

	number, _ := strconv.Atoi(os.Args[2])

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for randomNumber := r.Intn(number + 1); ; randomNumber = r.Intn(number + 1) {
		if verbose {
			fmt.Printf("%d ", randomNumber)
		}
		if randomNumber == number {
			fmt.Printf("🎉  YOU WIN!\n")
			break
		}
	}
}
