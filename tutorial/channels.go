package main

import "fmt"

func testchannels() {

	// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	messages := make(chan string, 2)
	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.

	messages <- "buffered"
	messages <- "channel"
	// Later we can receive these two values as usual.

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
