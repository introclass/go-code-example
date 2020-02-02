package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome server 01")
	})
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome server 02")
	})
	return e
}

func main() {
	g := errgroup.Group{}

	server01 := &http.Server{
		Addr:    ":8080",
		Handler: router01(),
	}

	server02 := &http.Server{
		Addr:    ":8081",
		Handler: router02(),
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
