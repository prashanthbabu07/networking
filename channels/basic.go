package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "ping2"
	}()

	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(<-messages)

	twoMessages := make(chan string, 2)

	twoMessages <- "buffered"
	twoMessages <- "channel"

	fmt.Println(<-twoMessages)
	fmt.Println(<-twoMessages)

}
