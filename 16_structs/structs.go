package main

import (
	"fmt"
	"time"
)

//Order Struct
//Struct used as an object oriented programming in golang

// golang has a convention whenever you want to change or modify a structs value use star on struct as you can see this on changeStatus function but when you want extract or get a value you dont need to use the {{{star or should I say pointer}}}

// struct has been use heavily used in golang.....

type customer struct {
	name  string
	phone string
}

type order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time // nanosecond precision

	//struct embedding like inheritence in class

	customer
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	fmt.Println("Structs")

	// if you dont set any field, default value will be zero or empty string

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 52.00,
	// 	status: "received",
	// }

	// myOrder := newOrder("1", 30.5, "received")

	// myOrder.createAt = time.Now()

	// myOrder.changeStatus("confirmed")

	// fmt.Println("Order 1 Amount", myOrder.getAmount())

	// myOrder2 := order{
	// 	id:       "2",
	// 	amount:   100.00,
	// 	status:   "delivered",
	// 	createAt: time.Now(),
	// }

	// fmt.Println("Order Struct", myOrder)
	// fmt.Println("My Order 2 Struct", myOrder2)

	//For one time use struct
	language := struct {
		name   string
		isGood bool
	}{"Golang", true}

	fmt.Println("Language", language)

	newCustomer := customer{
		name:  "John Doe",
		phone: "123456789",
	}

	newOrder := order{
		id:       "1",
		amount:   354,
		status:   "received",
		customer: newCustomer,
	}

	fmt.Println(newOrder)

}
