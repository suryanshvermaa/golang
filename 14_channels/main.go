package main

import (
	"fmt"
)

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("prcessing num channel", num)
// 	}
// }

// func sum(numChan chan int, a int, b int) {
// 	ans := a + b
// 	numChan <- ans
// }

func task(done chan bool) {
	defer func() {
		done <- true
	}()

	fmt.Println("Processing task...")
}

func main() {
	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// 	time.Sleep(time.Second)
	// }
	// go sum(numChan, 10, 20)
	// fmt.Println(<-numChan) //blocking

	// done := make(chan bool)

	// go task(done)
	// <-done // block

	// buffered channel
	emailChan := make(chan string, 100)
	emailChan <- "suryansh1@gmail.com"
	emailChan <- "suryansh2@gmail.com"
	emailChan <- "suryansh3@gmail.com"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)
}
