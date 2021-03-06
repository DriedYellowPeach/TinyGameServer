/*
@Time    : 12/5/21 18:52
@Author  : nil
@File    : gate.go
*/

package module

import (
	"bufio"
	"net"
)

type Router struct {
	Modules []*Core
}

func (r *Router) Init() {
	r.Modules = []*Core{}
}

func (r *Router) Route(msg string) (*Core, string){
	//根据msg判断dst
	return r.Modules[0], "default"
}

type Gate struct {
	router *Router
	c *Core
}

func (g *Gate) Init() {
	g.c = NewCore()
	g.router = new(Router)
	g.router.Init()
}

func (g *Gate) Run(){
	g.c.Run()
}

func (g *Gate) SetDefault(core *Core) {
	g.router.Modules = append(g.router.Modules, core)
}

func (g *Gate) Route(msg string, a *Agent) {
	mod, msgHandler := g.router.Route(msg)
	g.c.Client.Bind(mod.Server)
	g.c.Client.AsyncCall(msgHandler, msg, a, nil)
}

type Agent struct{
	conn net.Conn
	gate *Gate
}

func NewAgent(conn net.Conn, gate *Gate) *Agent {
	a := new(Agent)
	a.gate = gate
	a.conn = conn
	return a
}

func (a *Agent) Run() {
	input := bufio.NewScanner(a.conn)
	for input.Scan() {
		line := input.Text()
		a.gate.Route(line, a)
	}
}

func (a *Agent) Write(msg string) {
	a.conn.Write([]byte(msg))
}