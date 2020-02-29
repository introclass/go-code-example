package main

import (
	"log"
	"net/http"
	"time"
)

type NewHandler struct{}

func (s *NewHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	// 根据 req 查找对应的 handler
	if req.URL.Path == "/hello2" {
		writer.Write([]byte("Hello2 World!"))
	}
}

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
	s.Handler = &NewHandler{}
	log.Fatal(s.ListenAndServe())
}

func main() {
	// $ curl 127.0.0.1:8082/hello2
	// Hello2 World!%
	go defServer(":8082")
	time.Sleep(1000 * time.Second)
}
