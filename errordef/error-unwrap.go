// Create: 2019/12/03 14:31:00 Change: 2019/12/03 14:42:06
// FileName: error-unwrap.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"errors"
)

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return e.Query + ": " + e.Err.Error()
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

func FunC() error {
	return errors.New("err in Function C")
}

func FunB() error {
	err := FunC()
	if err != nil {
		return &QueryError{Query: "query error", Err: err}
	}
	return nil
}

func main() {
	err := FunB()
	if err != nil {
		println(err.Error())
		if inter := errors.Unwrap(err); inter != nil {
			println(inter.Error())
		}
	}
}
