// Create: 2019/07/05 11:04:00 Change: 2019/07/05 11:07:18
// FileName: json_and_others.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

type A struct {
	A []string
}

func main() {
	a := A{}
	if a.A == nil {
		println("is nil")
	}
}
