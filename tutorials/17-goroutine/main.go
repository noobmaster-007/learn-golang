package main

import (
	"fmt"
	"time"
)

func main() {
	// make a channel with string type
	c := make(chan string)
	go count("Sheep", c)

	// When doing receiving, it block the main routine
	for msg := range c {
		fmt.Println(msg)
	}

	// create a buffered channel for 2 message
	c2 := make(chan string, 2)
	c2 <- "Hello"
	c2 <- "World"

	msg := <-c2
	fmt.Println(msg)

	msg = <-c2
	fmt.Println(msg)

	// go routing with select
	c3 := make(chan string)
	c4 := make(chan string)

	go func() {
		for {
			c3 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c4 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// receiving msg is blocking

	// for {
	// 	fmt.Println(<-c3)
	// 	fmt.Println(<-c4)
	// }

	// Use select to pick up msg when it is ready
	for {
		select {
		case msg1 := <-c3:
			fmt.Println(msg1)
		case msg2 := <-c4:
			fmt.Println(msg2)
		}
	}
}

func count(thing string, c chan string) {
	for i := 0; i < 5; i++ {

		// send message to channel
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	// close the channel, only the sender should close it.
	close(c)
}
