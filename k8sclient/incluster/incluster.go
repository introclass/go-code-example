package main

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		glog.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatal(err.Error())
		return
	}
	spaceList, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		glog.Fatal(err.Error())
	}
	for _, v := range spaceList.Items {
		fmt.Println("namesapce: %s\n", v.Name)
	}
	time.Sleep(1000 * time.Second)
}
