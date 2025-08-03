package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum = sum + val
	}
	return sum
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
