// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat: lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/golang/glog"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// create restClient from config
func CreateCoreRestClient(config *rest.Config) (*rest.RESTClient, error) {

	config.ContentConfig.GroupVersion = &core_v1.SchemeGroupVersion
	config.ContentConfig.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.APIPath = "/api"

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		return nil, err
	}

	if restClient == nil {
		return nil, errors.New("restclient1 is nil")
	}

	return restClient, nil
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

	//方法1： 直接创建 restClient
	restClient, err := CreateCoreRestClient(config)
	if err != nil {
		glog.Fatal(err.Error())
		return
	}

	// /<namespace>/<resource>/<name>
	// GET https://10.10.173.203/api/v1/namespaces/default/pods?limit=500
	// GET https://10.10.173.203/api/v1/namespaces/kube-system/pods?limit=500
	// GET https://10.10.173.203/api/v1/namespaces/kube-system/pods/kube-dns-5b54cf547c-jl4r9
	result := restClient.Get().
		Namespace("kube-system").
		Resource("pods").
		Name("kube-dns-5b54cf547c-jl4r9").
		Do()
	bytes, err := result.Raw()
	if err != nil {
		fmt.Printf("%s: %s\n", err.Error(), bytes)
	} else {
		fmt.Printf("%s\n", bytes)
	}

	//方法2：从 clientSet 中获取
	restInterface, err := GetCoreRestClient(config)
	if err != nil {
		glog.Fatal(err.Error())
		return
	}

	result = restInterface.Get().
		Namespace("kube-system").
		Resource("pods").
		Name("kube-dns-5b54cf547c-jl4r9").
		Do()
	bytes, err = result.Raw()
	if err != nil {
		fmt.Printf("%s: %s\n", err.Error(), bytes)
	} else {
		fmt.Printf("%s\n", bytes)
	}
}
