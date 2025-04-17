package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Calculating the investment")

	const inflationRate = 6.5

	var investmentAmount float64 = 1000
	var expectedRateOfReturn float64 = 5.5
	var years float64 = 10

	fmt.Println("Please enter the investment amount:")
	_, err := fmt.Scan(&investmentAmount)

	if err != nil {
		fmt.Println(err)
		return
	}

	futureValue := investmentAmount * math.Pow(1+expectedRateOfReturn/100, years)
	futureRealValue := futureValue * math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value is: %v \n", futureValue)
	fmt.Printf("Future real value is: %v \n", futureRealValue)

}
