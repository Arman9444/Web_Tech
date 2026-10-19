package main

import (
	"fmt"
	"time"
)

// func processNum(numChan chan int) {
// 	fmt.Println("Processing number ", <-numChan)
// }

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number ", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	summation := num1 + num2
	result <- summation
}

func emailSender(emailchan1 chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailchan1 {
		fmt.Println("Sending email to ", email)
		time.Sleep(time.Second)
	}
}
func main() {
	// messageChan := make(chan string)
	// messageChan <- "ping"
	// msg := <-messageChan //deadlock
	// fmt.Println(msg)

	// numChan := make(chan int)
	// go processNum(numChan)
	// numChan <- 5
	// time.Sleep(time.Second)

	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//blocking channel
	result := make(chan int)
	go sum(result, 4, 5)
	sum := <-result
	fmt.Println(sum)

	//buffer channel
	emailchan := make(chan string, 100)
	emailchan <- "1@gmail.com"
	emailchan <- "2@gmail.com"
	fmt.Println(<-emailchan)
	fmt.Println(<-emailchan)
	//here though 2 chanel were called at the same time it will not fall into deadlock
	//because the size were declared,,, if after 100 times more msg were sent then it will fall into deadlock

	emailchan1 := make(chan string, 100)
	done := make(chan bool)
	go emailSender(emailchan1, done)
	for i := 1; i <= 5; i++ {
		emailchan1 <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(emailchan1)
	<-done
}
