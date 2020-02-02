// Create: 2019/12/02 11:14:00 Change: 2019/12/02 11:47:44
// FileName: wrong_done.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"sync"
	"time"
)

var a string
var done bool
var once sync.Once

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	println(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	done = false
	a = "not hello, world"
	twoprint()
	time.Sleep(1 * time.Second)
}
