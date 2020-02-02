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

func isEmpty(m *sync.Map) bool {
	isEmpty := true
	m.Range(func(key, value interface{}) bool {
		isEmpty = false
		return false
	})
	return isEmpty
}

func main() {
	//	m1 := &sync.Map{}
	//
	//	go insert(m1)
	//	//	go delete(m1)
	//	go read(m1)
	//
	//	time.Sleep(5 * time.Second)
	//	m1.Range(printFunc)

	m2 := &sync.Map{}
	if isEmpty(m2) {
		println("1: is empty")
	} else {
		println("1: not empty")
	}

	m2.Store(1, 1)
	if isEmpty(m2) {
		println("2: is empty")
	} else {
		println("2: not empty")
	}

	m2.Delete(1)
	if isEmpty(m2) {
		println("2: is empty")
	} else {
		println("2: not empty")
	}
	time.Sleep(5 * time.Second)
}
