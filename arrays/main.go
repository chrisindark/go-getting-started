package main

import "fmt"

func main() {
	// default values for numbers in an array is 0
	var numbers [5]int
	fmt.Println(numbers)

	// default values for strings in an array is empty string
	var strings [5]string
	fmt.Println(strings)

	definedNumbers := []int{1, 2, 3, 4, 5}
	fmt.Println(definedNumbers)

	definedStrings := []string{"chris", "paul", "go", "gin"}
	fmt.Println(definedStrings)
}
