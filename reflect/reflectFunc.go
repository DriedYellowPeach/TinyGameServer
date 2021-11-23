/*
@Time    : 11/23/21 07:53
@Author  : nil
@File    : reflectFunc.go
*/

package main

import (
	"fmt"
	"reflect"
)

func hello() {
	fmt.Println("hello world")
}

func main() {
	h1 := hello
	fv := reflect.ValueOf(h1)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)
}

