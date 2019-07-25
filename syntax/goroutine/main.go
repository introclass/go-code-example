// Create: 2019/07/01 14:52:00 Change: 2019/07/01 15:06:04
// FileName: modelbind.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"time"
)

func SubRoute() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				println("1")
			}
		}
	}()
}

func main() {
	SubRoute()
	time.Sleep(10 * time.Second)
}
