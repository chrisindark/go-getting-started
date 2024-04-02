package main

import "fmt"

var a = 20
var b = &a

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
}
