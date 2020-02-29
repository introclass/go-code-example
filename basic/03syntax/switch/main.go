// Create: 2019/07/03 19:30:00 Change: 2019/07/08 17:48:03
// FileName: json_and_others.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

func main() {
	i := 1
	switch i {
	case 1:
		fallthrough
	case 2:
		println("ok")
	}

	str := "abc"
	switch str {
	case "123":
		println("123")
	case "abc":
		println("abc")
	}

	str1 := "abc"
	switch str {
	case "123":
		println("123")
	case str1:
		println("abc")
	}
}
