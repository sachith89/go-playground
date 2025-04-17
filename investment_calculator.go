package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")

	var investmentAmount, expectedRateOfReturn, years float64 = 1000, 5.5, 10

	futureValue := investmentAmount * math.Pow(1+expectedRateOfReturn/100, years)

	fmt.Println(futureValue)
}
