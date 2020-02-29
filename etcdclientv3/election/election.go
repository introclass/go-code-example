// Create: 2019/07/10 15:58:00 Change: 2019/07/11 11:37:16
// FileName: json_and_others.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"context"

	"github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
	"go.etcd.io/etcd/clientv3/concurrency"
)

func PrintGetRespValue(resp *clientv3.GetResponse) {
	for i, v := range resp.Kvs {
		fmt.Printf("%d: %s %s\n\t%s\n", i, v.Key, v.Value, v.String())
	}
}

func main() {
	var node string
	flag.StringVar(&node, "name", "node1", "node name for election")

	flag.Parse()

	electprefix := "/testelection"
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		glog.Fatal(err.Error())
	}

	defer func() {
		if err := cli.Close(); err != nil {
			glog.Error(err.Error())
		}
	}()

	var s *concurrency.Session
	s, err = concurrency.NewSession(cli, concurrency.WithTTL(1))
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	e := concurrency.NewElection(s, electprefix)

	// leader
	newleader := e.Observe(context.Background())
	go func() {
		for {
			select {
			case resp := <-newleader:
				fmt.Printf("New Election Result:\n")
				PrintGetRespValue(&resp)
			}
		}
	}()

	//竞选 Leader，直到成为 Leader 函数才返回
	if err = e.Campaign(context.Background(), node); err != nil {
		glog.Fatalf("Campaign() returned non nil err: %s", err)
	}
	fmt.Printf("I'm leader")

	time.Sleep(5 * time.Minute)
}
