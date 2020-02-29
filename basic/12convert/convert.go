package main

import (
	"fmt"
	"math"
)

func main() {
	var a int32 = 80
	var b int32 = 100
	c := int32(math.Floor(float64(b) * (float64(a) / 100)))
	fmt.Println("%d", c)

	a = -3
	fmt.Println("absoulte is: ", math.Abs(float64(a)))
}
