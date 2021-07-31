package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("Hi", c)

	
	// for{
	// 	msg, isOpen := <- c
	// 	fmt.Println(msg)
	// 	if !isOpen {
	// 		break
	// 	}
	// }

	//equivalent
	
	for msg := range c {
		fmt.Println(msg)
	}
}


func count(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		// fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}