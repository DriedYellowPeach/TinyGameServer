/*
@Time    : 12/5/21 18:01
@Author  : nil
@File    : core.go
*/

package module

import (
	"game/chaninvoke"
)

type Core struct {
	Client             *chaninvoke.Client
	Server             *chaninvoke.Server
}

func NewCore() *Core{
	c := new(Core)
	c.Server = chaninvoke.StartServer(10)
	c.Client = chaninvoke.StartClient(10, nil)
	return c
}

func (c *Core) Run() {
	for {
		select {
		case ri := <-c.Client.AsyncRetChan:
			//log.Println("cb")
			c.Client.Cb(ri)
		case ci := <-c.Server.CallChan:
			//log.Println("exec:%d, %d", len(c.Server.CallChan), cap(c.Server.CallChan))
			c.Server.Exec(ci)
		}
	}
}

