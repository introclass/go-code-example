package main

import (
	"fmt"
	"time"
)

func Write(num int, c chan int) {
	for {
		c <- num
	}
}

func main() {
	c := make(chan int)
	go Write(10, c)
	go Write(20, c)
	for {
		select {
		case v := <-c:
			fmt.Printf("receive %d\n", v)
			time.Sleep(2 * time.Second)
		}
	}
}
