package main

import (
	"fmt"
)

var name = "John Doe"

// Declare the structure
type Employee struct {
	fName, lName string
	age          int
}

func main() {
	var hello string = "Hiii"
	email := "john@gmail.com"
	fmt.Println(hello, name, email)

	// Multiline variable declaration
	var one, two, three, four = 101, 2054, 3984565, 4798
	fmt.Println(one, two, three, four)

	// constant declaration
	const user = "Admin" //can not be changed
	fmt.Println(user)

	// Boolean declaration
	var isFlutterGood bool = true
	isGoLangTough := false
	fmt.Println(!isGoLangTough)
	fmt.Println("The Flutter is good for cross platform applications:", isFlutterGood)

	// Arithmetic Operators
	fmt.Println(2 + 2)
	fmt.Println(8 - 2)
	fmt.Println(4 / 2)
	fmt.Println(2 * 2)
	fmt.Println(9 % 2)

	// Named Structs
	fmt.Println("Struct in Go...!")
	emp1 := Employee{
		fName: "Vel",
		lName: "Murugan",
		age:   23,
	}
	emp2 := Employee{
		"Varan",
		"Raghu",
		22,
	}
	fmt.Println(emp1)
	fmt.Println(emp2)

	// Anonymous structs
	userData := struct{ fname, lname string }{
		"Vel", "murugan",
	}
	fmt.Println(userData)

	// Function
	total := calcBill(12, 78)
	fmt.Println(total)

	// Method
	emp1.getUserData()
	emp2.getUserData()

}

// Function in Go
func calcBill(a, b int) int {
	return a + b
}

// Method
func (e Employee) getUserData() {
	println(e.fName, e.lName, "Age: ", e.age)
}
