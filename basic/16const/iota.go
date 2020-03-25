package main

import "fmt"

const (
	ErrorCodeConfig = 60000300 + iota
	ErrorCodeConfigList
	ErrorCodeConfigList2
)

const (
	a = 100 * iota //0
	b = 100 * iota
	c = 100 * iota
)

const (
	A = c + iota
	B
	C
)

func main() {
	fmt.Println(ErrorCodeConfigList)
	fmt.Println(ErrorCodeConfigList2)
	fmt.Printf("c is: %d\n", c)
	fmt.Printf("C is: %d\n", C)
}
