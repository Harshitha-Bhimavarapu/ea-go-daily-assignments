package main

import (
	"errors"
	"fmt"
	"math"
)

var errorOnDividedByZero = errors.New("divided by Zero is not allowed")

func addTwoNumbers(number1 int, number2 int) int {
	return number1 + number2
}

func subtractTwoNumbers(number1 int, number2 int) int {
	return number1 - number2
}

func multiplyTwoNumbers(number1 int, number2 int) int {
	return number1 * number2
}

func divisionOfTwoNumbers(number1, number2 float64) (float64, error) {
	if number2 == 0 {
		return 0, errorOnDividedByZero
	}
	return number1 / number2, nil
}

func computeSine(radian float64) float64 {
	return math.Sin(radian)
}

func computeCosine(radian float64) float64 {
	return math.Cos(radian)
}

func computeTan(radian float64) float64 {
	return math.Tan(radian)
}

func computeSquareRoot(rootValue float64) float64 {
	return math.Sqrt(rootValue)
}

func main() {
	number1, number2 := 1, 2
	fmt.Println(addTwoNumbers(number1, number2))
	fmt.Println(subtractTwoNumbers(number1, number2))
	fmt.Println(multiplyTwoNumbers(number1, number2))
	fmt.Println(divisionOfTwoNumbers(100, 2))
	fmt.Println(computeSine(0))
	fmt.Println(computeCosine(0))
	fmt.Println(computeTan(0.785))
	fmt.Println(computeSquareRoot(25))
}
