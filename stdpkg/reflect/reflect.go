// Create: 2019/07/01 10:48:00 Change: 2019/07/01 14:52:31
// FileName: modelbind.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"reflect"
)

type Struct1 struct{}

func subfunc1(obj interface{}) {
	fmt.Println(reflect.TypeOf(obj).Name())
}

func main() {
	s1 := Struct1{}
	fmt.Println(reflect.TypeOf(s1).Name())
	subfunc1(s1)
	fmt.Println("abc:", reflect.TypeOf(subfunc1).Name())
}
