package main

import "fmt"

func main() {
	accountBalance := 100
	accountCredit := 200

	if accountBalance+accountCredit > 0 {
		fmt.Println("you have money to spend")
	} else {
		fmt.Println("you have no money left to spend")
	}

	score := 99
	scoreO := 90
	scoreE := 80
	scoreA := 70
	scoreB := 60
	scoreC := 50
	scoreD := 40
	scoreF := 0

	if score >= scoreO {
		fmt.Println("Outstanding")
	} else if score >= scoreE && score < scoreO {
		fmt.Println("Excellent")
	} else if score >= scoreA && score < scoreE {
		fmt.Println("Very Good")
	} else if score >= scoreB && score < scoreA {
		fmt.Println("Good")
	} else if score >= scoreC && score < scoreB {
		fmt.Println("Satisfactory")
	} else if score >= scoreD && score < scoreC {
		fmt.Println("Poor")
	} else if score >= scoreF && score < scoreD {
		fmt.Println("Fail")
	}

	hasGas := true
	hasKeyInIgnition := true

	if hasGas && hasKeyInIgnition {
		fmt.Println("Can drive car")
	}

	hasBurger := true
	hasSandwich := false

	if hasBurger || hasSandwich {
		fmt.Println("Can eat")
	}

	if !hasSandwich {
		fmt.Println("No sandwiches, then I will starve, I only eat sandwiches")
	}
}
