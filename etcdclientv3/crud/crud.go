// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat: lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
	"go.etcd.io/etcd/clientv3"
)

func toString(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PrintGetRespValue(resp *clientv3.GetResponse) {
	for i, v := range resp.Kvs {
		fmt.Printf("%d: %s %s\n\t%s\n", i, v.Key, v.Value, v.String())
	}
}

// 写入ETCD
func PUT(cli *clientv3.Client, key, val string, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := cli.Put(ctx, key, val, opts...)
	cancel()
	if err != nil {
		return err
	}

	respStr, err := toString(resp)
	if err != nil {
		return err
	}

	fmt.Printf("PUT RESULT: %s\n", respStr)
	if resp.PrevKv != nil {
		fmt.Printf(">> prev key: %s value: %s\n", resp.PrevKv.Key, resp.PrevKv.Value)
	}
	fmt.Printf("\n")

	return nil
}

// 查询ETCD
func GET(cli *clientv3.Client, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := cli.Get(ctx, key, opts...)
	cancel()
	return resp, err
}

func WATCH(cli *clientv3.Client, key string, opts ...clientv3.OpOption) (clientv3.WatchChan, error) {
	watchChan := cli.Watch(context.Background(), key, opts...)
	return watchChan, nil
}

func main() {

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

	if err := PUT(cli, "/dir/sample_key", "sample_value123", clientv3.WithPrevKV()); err != nil {
		glog.Errorf(err.Error())
	}

	if resp, err := GET(cli, "/dir/sample_key"); err != nil {
		glog.Errorf(err.Error())
	} else {
		respStr, _ := toString(resp)
		fmt.Printf("GET RESULT: %s\n", respStr)
		fmt.Printf("> count: %d  more: %v\n", resp.Count, resp.More)
		for i, v := range resp.Kvs {
			fmt.Printf(">> %d key: %s value: %s\n", i, v.Key, v.Value)
		}
		fmt.Printf("\n")
	}

	if resp, err := GET(cli, "/dir", clientv3.WithPrefix()); err != nil {
		glog.Errorf(err.Error())
	} else {
		PrintGetRespValue(resp)
	}

	if watchChan, err := WATCH(cli, "/dir", clientv3.WithPrefix(), clientv3.WithPrevKV()); err != nil {
		glog.Errorf(err.Error())
	} else {
		for {
			wr := <-watchChan
			if wrStr, err := toString(wr); err != nil {
				glog.Error(err.Error())
			} else {
				fmt.Printf("WATCH RESULT: %s\n", wrStr)
				for i, v := range wr.Events {
					if v.Kv != nil {
						fmt.Printf("> %d type: %d  key: %s value: %s ", i, v.Type, v.Kv.Key, v.Kv.Value)
					}
					if v.PrevKv != nil {
						fmt.Printf("pre value: %s\n", v.PrevKv.Value)
					}
				}
				fmt.Printf("\n")

			}
		}
	}
}
