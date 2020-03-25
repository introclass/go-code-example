package main

import "fmt"

//func illegeDiv() {
//	fmt.Printf("%d", 1/0)
//}
func illegeFloatDiv(a, b float64) {
	fmt.Println("%d", a/b)
}

func float64Euqal(a, b float64) {
	if a == b {
		fmt.Println("float64Euqal")
	}
}

func main() {
	//illegeDiv()
	illegeFloatDiv(0, 0)
	float64Euqal(0, 0)
}
