package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func test() {
}

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	fmt.Println("argument: ", arg1)
	fmt.Println("argument type: ", reflect.TypeOf(arg1))

	no1, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Println(err)
		fmt.Println("could not convert:", arg1)
	}

	no2, err := strconv.Atoi(arg2)
	if err != nil {
		fmt.Println(err)
		fmt.Println("could not convert:", arg2)
	}

	sum := add(no1, no2)
	fmt.Println("sum is: ", sum)

	var intStr string = "100"
	fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 8)    // becomes no 16 and int64
	tenBaseSixteenBitInt, _ := strconv.ParseInt(intStr, 10, 16) // no 100,  and int64
	fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
	fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))

	var noOfPlayers = 8
	str := strconv.Itoa(noOfPlayers)
	fmt.Println("str is: ", str)
	fmt.Println("type of str is: ", reflect.TypeOf(str))
}
