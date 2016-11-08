package main

import (
	"fmt"
	"math"
)

func SayHello(who string) {
	fmt.Printf("Hello %s\n", who)
}

func CalculateSin(x float64) float64 {
	return math.Sin(x)
}

// Increment next year ;-)
var CurrentYear = 2016
