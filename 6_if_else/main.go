package main

import "fmt"

func main() {
	age := 19
	// if age >= 18 {
	// 	fmt.Println("Person is adult")
	// } else {
	// 	fmt.Println("Person is not adult")
	// }
	if age >= 18 {
		fmt.Println("Person is adult")
	} else if age >= 12 {
		fmt.Println("Person is tennager")
	} else {
		fmt.Println("Person is not adult")
	}

	var role = "admin"
	var hasPermissions = true
	if role == "admin" && hasPermissions {
		fmt.Println("Yes")
	}
	if age := 15; age >= 18 {
		fmt.Println("Person is adult", age)
	} else if age >= 12 {
		fmt.Println("Person is tennation", age)
	}
	//go does not have ternary operator, you will have to use normal if else
}
