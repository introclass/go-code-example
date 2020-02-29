package main

import (
	"log"
	"net/http"
	"time"
)

func defServer(addr string) {
	s := &http.Server{
		Addr:              addr,
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	s.Handler = serveMux()
	log.Fatal(s.ListenAndServe())
}

func serveMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello3", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello3 World!"))
	})
	return mux
}

func main() {
	// $ curl 127.0.0.1:8083/hello3
	//   Hello3 World!
	go defServer(":8083")
	time.Sleep(1000 * time.Second)
}
