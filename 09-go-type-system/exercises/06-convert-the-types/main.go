// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert the Types
//
//  Convert the variables to appropriate types.
//
// EXPECTED OUTPUT
//  325.5 299.5
// ---------------------------------------------------------

func main() {
	// DONT TOUCH THIS:
	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	// DONT TOUCH THIS:
	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	// ----------------------------------------
	// FIX THE CODE BELOW:
	// You should solve it only by using conversions.
	// Do not change the code in any other way.

	celsius *= Celsius(celsiusDegree * Temperature(factor))
	fahr *= Fahrenheit(fahrDegree * Temperature(factor))

	// ----------------------------------------
	// DONT TOUCH THIS
	fmt.Println(celsius, fahr)

	// YOU MAY REMOVE THESE WHEN YOU'RE DONE
	_, _, _ = celsiusDegree, fahrDegree, factor
}
