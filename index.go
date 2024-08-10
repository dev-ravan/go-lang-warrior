package main

import (
	"fmt"
	"strings"
)

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

	// String concatenation
	var greetings = "Hi " + " my name is" + " Ravan"
	fmt.Println(greetings)

	// Query strings
	// fmt.Printf("%c", greetings[7])
	fmt.Println(len(greetings))

	// Strings pkg
	fmt.Println(strings.ToLower(greetings))
	fmt.Println(strings.ToUpper(greetings))

	// Conditional Statement
	emailId := "velmurugan1211r@gmail.com"

	if strings.Contains(emailId, "@gmail.com") {
		fmt.Println("It's a valid mail id..!")
	} else {
		fmt.Println("Email id not valid")
	}

	// Lists
	fruits := []string{"Apple", "Banana", "Orange", "Mango"}
	fmt.Println(fruits)

	// For loop with continue
	for i := 0; i < len(fruits); i++ {
		if fruits[i] == "Mango" {
			continue
		}
		fmt.Println(fruits[i])
	}
	// For loop with break
	for i := 0; i < len(fruits); i++ {
		if fruits[i] == "Banana" {
			break
		}
		fmt.Println(fruits[i])
	}

	// We don't have While loop in Go.
	// But we can use for loop as while loop.
	i := 1
	for i <= 5 {
		fmt.Println("Hello world ", i)
		i++
	}

	// Switch statement
	// roleId := 7

	// switch roleId {
	// case 0:
	// 	println("Owner")
	// case 1:
	// 	println("Manager")
	// case 2:
	// 	println("Employee")
	// case 3:
	// 	println("Human Resource")
	// case 4:
	// 	println("Team Lead")
	// default:
	// 	println("Nothing like you!")
	// }

	// Arrays
	var myList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(cap(myList))

	// Maps
	userInfo := map[string]any{
		"id":     "001",
		"name":   "Velmurugan",
		"email":  "vel@gmail.com",
		"mobile": 9361797872,
	}

	for key, val := range userInfo {
		fmt.Println(key, val)
	}

}
