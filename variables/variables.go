package variables

import "fmt"

var firstName string = "Chris"

var (
	lastName string = "Paul"
	age      uint8  = 29
)

func test() {
	nationality := "indian"
	fmt.Printf("%s's nationality is %s \n", firstName, nationality)
}

func main() {
	fmt.Println("variables", firstName, lastName, age)

	fmt.Printf("%s %s is %d years old \n", firstName, lastName, age)

	test()
}
