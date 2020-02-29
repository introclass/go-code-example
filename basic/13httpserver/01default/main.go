package main

import (
	"log"
	"net/http"
	"time"
)

func defaultServer(addr string) {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	// $ curl 127.0.0.1:8081/hello
	// Hello World!%
	go defaultServer(":8081")
	time.Sleep(1000 * time.Second)
}
