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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Pick a number")
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	randGen := rand.New(rand.NewSource((time.Now().UnixNano())))

	if randomNumber := randGen.Int(); guess == randomNumber {
		fmt.Println("Bonus message")
	} else {
		fmt.Println("Ohh ...")
	}
}
