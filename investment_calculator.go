package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")

	investmentAmount := 1000.0
	expectedRateOfReturn := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1+expectedRateOfReturn/100, years)

	fmt.Println(futureValue)
}
