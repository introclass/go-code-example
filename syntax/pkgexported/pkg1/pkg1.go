// Create: 2019/06/12 15:33:00 Change: 2019/06/12 15:35:53
// FileName: pkg1.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package pkg1

//[Exported identifiers](https://golang.org/ref/spec#Exported_identifiers)
//
//An identifier may be exported to permit access to it from another package. An identifier is exported if both:
//
//1. the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu");
//2. the identifier is declared in the package block or it is a field name or method name.
//
//All other identifiers are not exported.

//导出的变量
var ExportedVar ExportedStruct

//不导出的变量
var unexportedVar unexportedStruct

//导出的函数
func ExportedFunc() {
	println("exported func")
}

//不导出的函数
func unexportedFunc() {
	println("unexported func")
}

//导出的 Struct 类型
type ExportedStruct struct {
	unexportedField string //未导出的 field
	ExportedField   string //导出的 field
}

//不导出的 Struct 类型
type unexportedStruct struct {
	unexportedField string //未导出的 field
	ExportedField   string //导出的 field
}

//导出的 Struct 的方法
func (s ExportedStruct) ExportedMethod() {
	println("exported method")
}

//不导出的 Struct 的方法
func (s ExportedStruct) unexportedMethod() {
	println("unexported method")
}

func ReturnUnexporteFunc() func() {
	return unexportedFunc
}

func init() {
	// 在 Package 内部，导出/未导出的变量、导出/未导出的 Struct 类型都可以使用
	ExportedVar = ExportedStruct{
		unexportedField: "unexported field",
		ExportedField:   "exported field",
	}

	unexportedVar = unexportedStruct{
		unexportedField: "unexported field",
		ExportedField:   "exported field",
	}

	// 未导出的 Struct 的未导出的 filed 可以使用
	unexportedVar.unexportedField = "unexported field"
}
