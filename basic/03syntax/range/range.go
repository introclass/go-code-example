package main

import "fmt"

var List []string

func init() {
	List = []string{
		"a",
		"b",
		"c",
	}
}

func main() {
	list := make([]int, 10)
	for i, v := range list {
		copyv := v
		fmt.Printf("i=%d v=%d copyv=%d\n", i, &v, &copyv)
	}

	for i, v := range List {
		copyv := v
		fmt.Printf("i=%d v=%s copyv=%d\n", i, v, &copyv)
	}
}
