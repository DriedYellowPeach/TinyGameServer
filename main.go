/*
@Time    : 12/5/21 20:53
@Author  : nil
@File    : main.go
*/

package main

import (
	"game/module"
	"game/network"
)

func main() {
	hello := new(module.Hello)
	hello.Init()
	go hello.Run()

	g := new(module.Gate)
	g.Init()
	g.SetDefault(hello.C)
	go g.Run()

	server := new(network.TCPServer)
	server.Addr = "127.0.0.1:8000"
	server.Init(g)

	server.Run()


}
