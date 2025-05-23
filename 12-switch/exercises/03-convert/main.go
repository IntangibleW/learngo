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
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

func main() {
	args := os.Args

	switch {
	case len(args) != 3:
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	switch u {
	case user:
		switch p {
		case pass:
			fmt.Printf(accessOK, u)
		default:
			fmt.Printf(errPwd, u)
		}
	case user2:
		switch p {
		case pass2:
			fmt.Printf(accessOK, u)
		default:
			fmt.Printf(errPwd, u)
		}
	default:
		fmt.Printf(errUser, u)
	}
}
