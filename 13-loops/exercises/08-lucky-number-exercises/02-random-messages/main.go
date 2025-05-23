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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

var (
	wonMessages = [...]string{
		"YOU WON",
		"YOU'RE AWESOME",
	}
	lostMessages = [...]string{
		"LOSER!",
		"YOU LOST. TRY AGAIN?",
	}
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me an argument")
		return
	}

	arg1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	Random := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch Random.Int() {
	case arg1:
		fmt.Println(wonMessages[Random.Intn(len(wonMessages))])
	default:
		fmt.Println(lostMessages[Random.Intn(len(lostMessages))])
	}
}
