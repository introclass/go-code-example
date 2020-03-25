package main

import (
	"fmt"
	"time"
)

func UnixSecond() {
	fmt.Printf("Unix Seconds: %d\n", time.Now().Unix())
}

func main() {
	UnixSecond()
}
