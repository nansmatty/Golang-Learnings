package main

import (
	"fmt"
	"time"
)

// If you want to receive or send data to one go routine to another go routine then channels helps us to create a path for the communication
// channels are blocking until the other side of channel path ready to receive

// This example is for sending
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// This example is for receiving
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// Making a mimic of wait-group which we learn on goroutines ............ goroutines synchronizer
// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Processing...")
// }

// Simple queue system using buffer channel
// With adding this arrow <- symbol we can make it more type safe adding arrow symbol before chan word make it only receiver in this function similarly adding the arrow after chan keywork make it only sender
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending Email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Channels")

	// messageChannel := make(chan string)
	// messageChannel <- "ping"
	// msg := <-messageChannel
	// fmt.Println(msg)

	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result
	// fmt.Println(res)

	// done := make(chan bool)
	// go task(done)
	// <-done //

	//Buffered Channel is not Blocking until the specified space is full and we sending more data
	// emailChan := make(chan string, 100)
	// done := make(chan bool)
	// go emailSender(emailChan, done)
	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("Done Sending...")
	// This is an example of closing the channel which is required to channel not to stuck in dedad lock
	// close(emailChan)
	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received Data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received Data from chan2", chan2Val)
		}
	}

}
