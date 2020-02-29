// Create: 2019/06/12 15:30:00 Change: 2019/06/27 11:01:23
// FileName: json_and_others.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import "lijiaocn.com/gocode/pkgexported/pkg1"

func main() {
	// 导出的 Struct 类型可用
	println("1")
	var v1 pkg1.ExportedStruct
	println(v1.ExportedField)
	// 未导出的 Struct 类型不可用
	// var v2 pkg1.unexportedStruct // wrong

	// 导出的变量可用
	println("2")
	v1 = pkg1.ExportedVar
	// 未导出的变量不可用
	// v2 = pkg1.unexportedVar

	// 导出的变量的导出的 filed 可以访问
	println("3")
	println(pkg1.ExportedVar.ExportedField)
	// 导出的变量的未导出的 filed 不可以访问
	// println(pkg1.ExportedVar.unexportedField) //wrong

	// 导出的变量的导出的方法可用
	println("4")
	pkg1.ExportedVar.ExportedMethod()
	// 导出的变量的未导出的方法不可用
	// pkg1.ExportedVar.unexportedMethod()

	// 导出的函数可用
	println("5")
	pkg1.ExportedFunc()
	// 未导出的函数不可用
	// pkg1.unexportedFunc()

	// 未导出的函数可以通过导出的函数获得
	println("6")
	unFunc := pkg1.ReturnUnexporteFunc()
	unFunc()
}
