package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	i := 0
	for i < 10 {
		fmt.Println("i: ", i)
		i++
	}

	arr := []string{"arg1", "arg2", "arg3"}
	for i, s := range arr {
		fmt.Printf("index: %d, item: %s \n", i, s)
	}

	arr1 := []int{-1, 2}
	for i := 0; i < 2; i++ {
		fmt.Println(arr1[i])
		if arr1[i] < 0 {
			break
		}
	}

	arr2 := []int{-1, 2, -1, 3}
	for i := 0; i < 4; i++ {
		if arr2[i] < 0 {
			continue
		}
		fmt.Println(arr2[i])
	}
}
