package main

import "fmt"

const age = 60

func main() {
	const name string = "golang"
	const age int = 30
	fmt.Println(name)
	fmt.Println(age)
	const (
		port = 500
		host = "localhost"
	)
	fmt.Println(port, host)
}
