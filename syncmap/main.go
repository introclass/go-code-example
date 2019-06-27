// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat: lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"math/rand"
	"sync"
	"time"
)

type Var struct {
	Value int
}

const NUM = 100

func insert(m *sync.Map) {
	for i := 1; i < NUM; i++ {
		key := i
		v := Var{Value: i}
		println("insert key: ", key)
		m.Store(i, v)
	}
}

func read(m *sync.Map) {
	for i := 1; i < NUM; i++ {
		key := i
		if v, ok := m.Load(i); ok {
			if value, ok := v.(Var); ok {
				println("read key: ", key, " value is: ", value.Value)
			} else {
				println("value's type is wrong")
			}
		} else {
			println("not found value for key: ", key)
		}
	}
}

func delete(m *sync.Map) {
	for i := 1; i < NUM/3; i++ {
		random := rand.Int()
		key := random % NUM
		println("delete key: ", key)
		m.Delete(key)
	}
}

func printFunc(key, value interface{}) bool {
	k, ok := key.(int)
	if !ok {
		println("key's type is wrong")
	}
	v, ok := value.(Var)
	if !ok {
		println("value's type is wrong")
	}
	println("range key is ", k, " value is ", v.Value)
	return true
}

func main() {
	m := &sync.Map{}

	go insert(m)
	//	go delete(m)
	go read(m)

	time.Sleep(5 * time.Second)
	m.Range(printFunc)
}
