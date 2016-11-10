package main

import (
	"fmt"
	"math"
)


// SayHello says hello :)
func SayHello(who string) {
	fmt.Printf("Hello %s\n", who)
}

// CalculateSin calculates sin for given x
func CalculateSin(x float64) float64 {
	return math.Sin(x)
}

// CurrentYear should be incremented next year ;-)
var CurrentYear = 2016
