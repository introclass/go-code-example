// Create: 2018/08/28 19:03:00 Change: 2019/07/08 17:20:13
// FileName: modelbind.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"time"
)

type Child func(input string)
type Child2 func(input string, c chan string)

func parent2(c chan string) Child2 {
	f := func(str string, c chan string) {
		for {
			c <- str
			time.Sleep(1 * time.Second)
		}
	}
	return f
}

func parent(c chan<- string) Child {
	f := func(str string) {
		for {
			c <- str
			time.Sleep(1 * time.Second)
		}
	}
	return f
}

func main() {

	c := make(chan string, 10)
	child := parent(c)
	go child("child1")

	//	child2 := parent2(c)
	//	go child2("child1", c)

	time.Sleep(5 * time.Second)
	println(len(c))

	for {
		select {
		case v := <-c:
			println(v)
		}
	}
}
