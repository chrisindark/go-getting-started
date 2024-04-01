package main

import "fmt"

func main() {
	// Define a map with string keys and int values
	var ages map[string]int

	// Initialize the map using make
	ages = make(map[string]int)

	// Add key-value pairs to the map
	ages["alice"] = 30
	ages["bob"] = 35
	ages["charlie"] = 40

	fmt.Println(ages)

	// Access values from the map
	fmt.Println("alice's age:", ages["alice"])
	fmt.Println("bob's age:", ages["bob"])
	fmt.Println("charlie's age:", ages["charlie"])

	// Check if a key exists in the map
	if _, value := ages["david"]; value {
		fmt.Println("david's age:", ages["david"])
	} else {
		fmt.Println("david's age does not exist in the map")
	}

	// Delete a key from the map
	delete(ages, "bob")
	fmt.Println("bob's age after deletion:", ages["bob"])

	fmt.Println(ages)
}
