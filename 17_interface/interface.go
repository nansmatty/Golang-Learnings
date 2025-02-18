package main

import (
	"fmt"
)

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

// while doing this we violated the solid principles open cloase method ................need to learn Coder's gyan had free in youtube

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw.pay(amount)

	// Modification aftersolid principle violation
	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	//logic to make payment
	fmt.Println("Making payment using stripe", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	//logic to make payment
	fmt.Println("Making payment using fake gateway for testing purpose", amount)
}

func main() {
	fmt.Println("Interface")

	// stripePaymentGw := stripe{}
	// fakePaymentGw := fakePayment{}
	razorpayPaymentGw := razorpay{}

	newPayment := payment{
		gateway: razorpayPaymentGw,
	}
	newPayment.makePayment(100)
}
