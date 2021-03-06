/*
@Time    : 11/21/21 20:06
@Author  : nil
@File    : chanInvoke.go
*/

package chaninvoke

import (
	"fmt"
	"reflect"
)

type Server struct {
	funcSet map[interface{}]interface{}
	CallChan chan * ArgPack
}

type Client struct {
	RetChan chan * RetPack
	AsyncRetChan chan * RetPack
	s *Server
}

type ArgPack struct {
	f interface{}
	args []interface{}
	RetChan chan * RetPack
	cb interface{}
}

type RetPack struct {
	ret interface{}
	err error
	cb interface{}
}

// StartServer
// l is the length of the call channel
func StartServer(l int) * Server {
	s := new(Server)
	s.funcSet = make(map[interface{}]interface{})
	s.CallChan = make(chan * ArgPack, l)
	return s
}

// todo: type assert
func (s *Server) Register(id interface{}, f interface{}) {
	s.funcSet[id] = f
}

func (s *Server) ret(ap *ArgPack, rp *RetPack) {
	if ap.RetChan == nil {
		return
	}
	// cb from the arg packet stay the same,
	//and cb will be call by corresponding client
	rp.cb = ap.cb
	ap.RetChan <- rp
}

// todo: * f exception handling
//       * f type assertion
//       * f is nil function, error handling
// 		 * 没有处理空返回值的情况
func (s *Server) Exec(arg *ArgPack) (err error) {
	//result := arg.f.(func(...interface{}) interface{})(arg.args...)

	//interface to type func
	fv := reflect.ValueOf(arg.f)

	// prepare the params, args is []interface{}, params is []reflect.Value
	var params []reflect.Value
	if arg.args == nil {
		params = nil
	} else if len(arg.args) == 0 {
		params = make([]reflect.Value, 0)
	} else {
		params = make([]reflect.Value, len(arg.args))
		for i := 0 ; i < len(arg.args) ; i++ {
			params[i] = reflect.ValueOf(arg.args[i])
		}
	}
	result := fv.Call(params)
	//如果result为nil, 这里会报错
	if len(result) == 0 {
		s.ret(arg, &RetPack{})
	} else {
		s.ret(arg, &RetPack{ret : result[0].Interface()})
	}
	return
}

func (s *Server) Call(id interface{}, args... interface{}){

}

func StartClient(l int, s *Server) *Client {
	c := new(Client)
	c.RetChan = make(chan *RetPack, l)
	c.AsyncRetChan = make(chan *RetPack, l)
	c.s = s
	return c
}

func (c *Client) Bind(s * Server) {
	c.s = s
}

func (c *Client) call(arg *ArgPack) {
	c.s.CallChan <- arg
}

// todo: async call
// this is sync call
func (c *Client) Call(id interface{}, args...interface{}) interface{} {
	f := c.s.funcSet[id]
	c.call(&ArgPack{
		f: f,
		args: args,
		RetChan: c.RetChan,
	})

	rp := <- c.RetChan
	return rp.ret
}

// AsyncCall the last of args is the cb
// 如果不需要cb, 那么最后一个参数一定要是nil
func (c *Client) AsyncCall(id interface{}, args...interface{}) interface {} {
	f := c.s.funcSet[id]
	l := len(args)
	//fmt.Printf("len: %d", l)

	c.call(&ArgPack{
		f:f,
		args:args[:l-1],
		RetChan: c.AsyncRetChan,
		cb: args[l-1],
	})
	return nil
}

func (c *Client) Cb(rp * RetPack) {
	//fmt.Println("here %v", rp.cb)
	if rp.cb == nil {
		//fmt.Println("cb is nil")
		return
	} else {
		fmt.Println("not nil")
		fv := reflect.ValueOf(rp.cb)
		params := reflect.ValueOf(rp.ret)
		fv.Call([]reflect.Value{params})
	}
}