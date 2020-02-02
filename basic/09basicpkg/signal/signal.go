// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat: lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func sigusr1() {
	fmt.Printf("Recevie SIGUSR1")
}

func main() {

	//Set signal
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGUSR1)
	for {
		select {
		case s := <-signalChan:
			switch s {
			case syscall.SIGQUIT:
				fallthrough
			case syscall.SIGKILL:
				fallthrough
			case syscall.SIGTERM:
				os.Exit(1)
			case syscall.SIGUSR1:
				sigusr1()
			default:
				continue
			}
		}
	}
}
