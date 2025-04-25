package main

import (
	"fmt"
	"time"
)

func main() {
	i := 4
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("Invalid")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday, time.Friday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("workday")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its's integer")
		case string:
			fmt.Println("it's string")
		default:
			fmt.Println("other", t)
		}
	}
	whoAmI("golang")
	whoAmI(true)
	whoAmI(55)
}
