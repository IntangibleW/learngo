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
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Give me 2 args")
		return
	}

	arg1, err1 := strconv.ParseUint(os.Args[1], 10, 64)
	arg2, err2 := strconv.ParseUint(os.Args[2], 10, 64)

	if err1 != nil {
		fmt.Println(err1)
		return
	} else if err2 != nil {
		fmt.Println(err2)
		return
	}

	Random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var randomNumber int
	if arg1 > arg2 {
		randomNumber = Random.Intn(int(arg1 + 1))
	} else {
		randomNumber = Random.Intn(int(arg2 + 1))
	}
	switch uint64(randomNumber) {
	case arg1, arg2:
		fmt.Println("Won")
	default:
		fmt.Println("Lose")
	}
}
