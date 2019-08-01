// Create: 2019/07/31 11:56:00 Change: 2019/07/31 11:57:03
// FileName: file.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"os"

	"github.com/golang/glog"
)

func main() {
	f, err := os.OpenFile("/tmp/go-open-file", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		glog.Fatal(err.Error())
		return
	}
	n, err := f.WriteString("hello")
	if err != nil {
		glog.Fatal(err)
	} else {
		fmt.Printf("Write %d\n", n)
	}
}
