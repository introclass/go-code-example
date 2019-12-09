// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat: lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"time"
)

type Node struct {
	Cluster string
	ID      string
	Time    time.Second
}

func main() {
	nodes := make([]Node, 0)
	nodes = append(nodes,
		Node{
			Cluster: "cluster",
			ID:      "1",
			Time:    5 * time.Second,
		},
		Node{
			Cluster: "cluster",
			ID:      "2",
		},
		Node{
			Cluster: "cluster",
			ID:      "3",
		},
	)

	str, err := json.Marshal(nodes)
	if err != nil {
		glog.Fatal(err)
	}
	fmt.Printf("%s\n", str)

	var nodes2 []Node
	if err := json.Unmarshal(str, &nodes2); err != nil {
		glog.Fatal(err)
	} else {
		for _, i := range nodes2 {
			fmt.Printf("%s %s\n", i.Cluster, i.ID)
		}
	}
}
