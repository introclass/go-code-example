// Create: 2019/07/08 17:20:00 Change: 2019/07/08 17:28:50
// FileName: json_and_others.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"errors"
	"fmt"
)

func send(errChan chan<- error) {
	for i := 0; i < 3; i++ {
		errChan <- errors.New("hello")
	}
}

func main() {
	c := make(chan error)
	go send(c)
	for {
		err := <-c
		fmt.Println(err.Error())
	}
}
