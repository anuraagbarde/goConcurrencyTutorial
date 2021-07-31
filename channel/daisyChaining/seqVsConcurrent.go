package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for{
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	
	go func() {
		for{
			c2 <- "Every 2000ms"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	// the below code is blocked since it is sequential
	// for{
	// 	fmt.Println(<-c1)	
	// 	fmt.Println(<-c2)
	// }

	// the below code is un-blocked since it is concurrent, it does not wait for one to finish before starting the other
	for{
		select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
		}
	}

}
