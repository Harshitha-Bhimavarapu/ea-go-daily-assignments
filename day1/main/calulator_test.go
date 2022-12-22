package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	result := addTwoNumbers(1, 2)

	if result != 3 {
		t.Error("Addition failed")
	}
}

func TestSubtractTwoNumbers(t *testing.T) {
	result := subtractTwoNumbers(10, 6)

	if result != 4 {
		t.Fatalf("Subtraction failed")
	}
}

func TestMultiplyTwoNumbers(t *testing.T) {
	result := multiplyTwoNumbers(10, 6)

	if result != 60 {
		t.Fatalf("Multiplication failed")
	}
}

func TestSuccessfulDivisionOfTwoNumbers(t *testing.T) {
	result, _ := divisionOfTwoNumbers(10, 5)

	if result != 2 {
		t.Fatalf("Division Failed")
	}
}

func TestFailsOnDivisionByZero(t *testing.T) {
	_, err := divisionOfTwoNumbers(10, 0)

	if err != errorOnDividedByZero {
		t.Fatalf("Error divided by zero")
	}
}

func TestComputeSinValue(t *testing.T) {
	const Sin0 = 0

	result := computeSine(0)

	if result != Sin0 {
		t.Error("Incorrect sine value")
	}
}

func TestComputeCosineValue(t *testing.T) {
	const Cosine0 = 1

	result := computeCosine(0)

	if result != Cosine0 {
		t.Error("Incorrect cosine value")
	}
}

func TestComputeTanValue(t *testing.T) {
	const Tan45 = 0.9992039901050427

	result := computeTan(0.785)

	if result != Tan45 {
		t.Error("Incorrect tan value")
	}
}
