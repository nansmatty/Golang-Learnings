package main

import "fmt"

//customs types
// type OrderStatus int
type OrderStatus string

//Constants concept
// iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing Order status to", status)
}

func main() {
	fmt.Println("Enums")

	changeOrderStatus(Prepared)
}
