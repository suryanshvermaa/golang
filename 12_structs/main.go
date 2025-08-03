package main

// struct embeddings
import "fmt"

type Customer struct {
	name  string
	phone string
}
type Order struct {
	Id     string
	Amount int
	status string
	Customer
}

func main() {
	order := Order{
		status: "dilivered",
		Id:     "1",
		Amount: 2000,
		Customer: Customer{
			name:  "Suryansh",
			phone: "9580104754",
		},
	}
	order.Customer.phone = "9580104753"
	fmt.Println(order)
}
