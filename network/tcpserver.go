/*
@Time    : 12/5/21 18:59
@Author  : nil
@File    : tcpserver
*/

package network

import (
	"game/module"
	"log"
	"net"
)

type TCPServer struct {
	Addr string
	ln net.Listener
	gate *module.Gate
}

func (svr *TCPServer) Init(gate *module.Gate) {
	ln, err := net.Listen("tcp", svr.Addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	svr.ln = ln
	svr.gate = gate

}

func (svr *TCPServer) Run() {
	for {
		conn, err := svr.ln.Accept()
		if err != nil {
			log.Printf("server Run error: %v", err)
		}

		go handleConn(conn, svr.gate)
	}
}

func handleConn(conn net.Conn, gate *module.Gate) {
	agent := module.NewAgent(conn, gate)
	agent.Run()
}
