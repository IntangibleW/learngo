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
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

const usage = "Usage: [username] [password]"

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
	}
	user, password := os.Args[1], os.Args[2]
	if user != "jack" {
		fmt.Printf("Access denied for %q\n", os.Args[1])
	} else if password != "1888" {
		fmt.Printf("Invalid password for %q\n", os.Args[1])
	} else {
		fmt.Printf("Access granted to %q\n", os.Args[1])
	}
}
