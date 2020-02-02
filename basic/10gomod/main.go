package main

import "lijiaocn.com/go/example/mod"
import "lijiaocn.com/go/example/mod/v2"
//import "lijiaocn.com/go/example/mod/v3"

func main() {
	mod.Version()
	v2.Version()
//	v3.Version()
}