package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	s = append(s, 5)
	s = append(s, 10)
	s = append(s, 11)
	fmt.Println(s)
	s2 := make([]int, 2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 10)
	s2 = append(s2, 11)
	fmt.Println(s2)
}
