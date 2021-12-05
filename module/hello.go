/*
@Time    : 12/5/21 20:17
@Author  : nil
@File    : hello.go
*/

package module

import (
	"fmt"
	"log"
)

type Hello struct {
	C *Core
	name string
}

func (h *Hello) Init() {
	h.name = "Hello"
	h.C = NewCore()
	h.C.Server.Register("default", func(msg string, a *Agent){
		a.Write(fmt.Sprintf("Hello from %s Module, RCV:%s\n", h.name, msg))
		log.Printf("Hello from %v", a.conn.RemoteAddr())
	})
}

func (h *Hello) Run() {
	h.C.Run()
}




