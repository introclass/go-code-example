// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat: lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"k8s.io/apimachinery/pkg/fields"

	"k8s.io/client-go/tools/cache"

	"github.com/golang/glog"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func printEndpoint(ep *core_v1.Endpoints) {
	fmt.Printf("%s in %s:\n", ep.Name, ep.Namespace)
	for i, sub := range ep.Subsets {
		fmt.Printf("\t%d: %s\n", i, sub.String())
	}
}

// get restClient from clientset
func GetCoreRestClient(config *rest.Config) (rest.Interface, error) {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	if clientset == nil {
		return nil, errors.New("clientset is nil")
	}

	restClient := clientset.CoreV1().RESTClient()
	if restClient == nil {
		return nil, errors.New("restclient is nil")
	}

	return restClient, nil
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		glog.Fatal(err.Error())
	}

	file, err := filepath.Abs(home + "/.kube/config")
	if err != nil {
		glog.Fatal(err.Error())
	}

	config, err := clientcmd.BuildConfigFromFlags("", file)
	if err != nil {
		glog.Fatal(err.Error())
		return
	}

	restClient, err := GetCoreRestClient(config)
	if err != nil {
		glog.Fatal(err.Error())
		return
	}

	labels := make(map[string]string)
	selector := fields.SelectorFromSet(labels)
	listwatch := cache.NewListWatchFromClient(restClient, "endpoints", "", selector)

	handler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("Add endpoint:\n")
			if ep, ok := obj.(*core_v1.Endpoints); !ok {
				fmt.Printf("not endpoints\n")
			} else {
				printEndpoint(ep)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("Update endpoint:\n")
			if epOld, ok := oldObj.(*core_v1.Endpoints); !ok {
				fmt.Printf("not endpoints\n")
			} else {
				printEndpoint(epOld)
			}

			if epNew, ok := newObj.(*core_v1.Endpoints); !ok {
				fmt.Printf("not endpoints\n")
			} else {
				printEndpoint(epNew)
			}
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("Delete endpoint:\n")
			if ep, ok := obj.(*core_v1.Endpoints); !ok {
				fmt.Printf("not endpoint")
			} else {
				printEndpoint(ep)
			}
		},
	}

	stop := make(chan struct{})
	store, controller := cache.NewInformer(listwatch, &core_v1.Endpoints{}, 0*time.Second, handler)
	go controller.Run(stop)

	time.Sleep(5 * time.Second)
	fmt.Printf("keys\n")
	for _, key := range store.ListKeys() {
		fmt.Printf("key: %s\n", key)
	}

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
				var obj struct{}
				stop <- obj
			case syscall.SIGUSR1:
				for i, v := range store.ListKeys() {
					fmt.Printf("%d : %s\n", i, v)
				}
			default:
				continue
			}
		}
	}
}
