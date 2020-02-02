// Create: 2019/07/03 17:14:00 Change: 2019/07/03 19:30:32
// FileName: modelbind.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

func func2(eles *[]*int) {
	for i := 0; i < 10; i++ {
		j := i
		*eles = append(*eles, &j)
	}
}

func main() {
	var eles1 []*int
	for i := 0; i < 10; i++ {
		j := i
		eles1 = append(eles1, &j)
	}

	for _, v := range eles1 {
		println(v, ":", *v)
	}

	println("eles2")
	var eles2 []*int
	func2(&eles2)
	for _, v := range eles2 {
		println(v, ":", *v)
	}
}
