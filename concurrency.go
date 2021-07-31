package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")
	
}


func count(thing string){
	for i := 0; i < 1000000; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}