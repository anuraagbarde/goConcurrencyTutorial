package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)

	c <- "one"
	c <- "two"
	// c <- "Three" // this will cause a deadlock panic

	msg := <- c
	fmt.Println(msg)
	
	msg = <- c
	fmt.Println(msg)

}
