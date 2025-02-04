// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Swapper #2
//
//  1. Swap the values of `red` and `blue` variables
//
//  2. Print them
//
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------

func main() {

	red, blue := "red", "blue"
	red, blue = swap(red, blue)
	fmt.Println(red, blue)
}

func swap(a, b string) (string, string) {
	return b, a
}
