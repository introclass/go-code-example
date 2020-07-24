//create: 2017/11/15 16:05:17 change: 2017/11/17 09:58:07 author:lijiao
package main

import (
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

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

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatal(err.Error())
		return
	}

	if clientset == nil {
		glog.Fatal("clientset is nil")
		return
	}

	namespace := "kube-system"
	pods, err := clientset.CoreV1().Pods(namespace).List(v1.ListOptions{})
	if err != nil {
		glog.Fatal(err.Error())
		return
	}

	if pods == nil {
		fmt.Printf("There are no pods in namespace %s.\n", namespace)
	} else {
		fmt.Printf("There are %d pods in namespace %s.\n", len(pods.Items), namespace)
	}

	ns, err := clientset.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		glog.Fatal(err.Error())
		return
	}

	if ns == nil {
		fmt.Printf("There are no namespaces.\n")
	} else {
		fmt.Printf("There are %d namespaces.\n", len(ns.Items))
	}

	namespace = "devops-selftest-1"
	pod, err := clientset.CoreV1().Pods(namespace).Get("nginx-996b77c56-mmws6", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	for i, v := range pod.Spec.Containers {
		fmt.Printf("%d: %s\n", i, v.Name)
		for ii, vv := range v.Resources.Requests {
			fmt.Printf("\trequest: %s  %.3f\n", ii, float64(vv.MilliValue())/1000)
		}
		for ii, vv := range v.Resources.Limits {
			fmt.Printf("\tlimit: %s  %d\n", ii, vv.Value())
		}
	}
}
