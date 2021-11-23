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

func display(i, j int) {
	fmt.Println(i, "->", j)
}

func main() {
	h1 := hello
	fv := reflect.ValueOf(h1)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)


	params := make([]reflect.Value, 2)
	var f2 interface{}
	var i interface{}
	f2 = display
	fv2 := reflect.ValueOf(f2)
	i = 5
	params[0]= reflect.ValueOf(i)
	params[1] = reflect.ValueOf(i)
	fv2.Call(params)
}

