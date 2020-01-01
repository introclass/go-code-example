package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func Write(num int, c chan int) {
	for {
		c <- num
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

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
