package main

import "fmt"

var name = "John Doe"

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
}
