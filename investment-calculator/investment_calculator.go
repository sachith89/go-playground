package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	fmt.Println("Calculating the investment")

	var investmentAmount float64 = 1000
	var expectedRateOfReturn = 5.5
	var years float64

	fmt.Print("Please enter the investment amount: ")
	_, err := fmt.Scan(&investmentAmount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Please enter the no of years: ")
	_, _ = fmt.Scan(&years)

	futureValue, futureRealValue := calculateInvestment(investmentAmount, expectedRateOfReturn, years, inflationRate)
	output(futureValue, futureRealValue)
}

func calculateInvestment(investmentAmount, expectedRateOfReturn, years, inflationRate float64) (futureValue float64, futureRealValue float64) {

	futureValue = investmentAmount * math.Pow(1+expectedRateOfReturn/100, years)
	futureRealValue = futureValue * math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
}

func output(futureValue, futureRealValue float64) {
	fmt.Printf("Future value is: %v \n", futureValue)
	fmt.Printf("Future real value is: %v \n", futureRealValue)
}
