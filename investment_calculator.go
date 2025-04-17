package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")

	investmentAmount, expectedRateOfReturn, years := 1000.0, 5.5, 10.0

	futureValue := investmentAmount * math.Pow(1+expectedRateOfReturn/100, years)

	fmt.Println(futureValue)
}
