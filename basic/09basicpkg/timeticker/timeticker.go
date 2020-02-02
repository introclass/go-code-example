// Create: 2019/07/01 14:14:00 Change: 2019/07/01 14:30:22
// FileName: modelbind.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			println("1")
		}
	}
}
