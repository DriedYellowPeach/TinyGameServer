/*
@Time    : 11/21/21 20:06
@Author  : nil
@File    : chanInvoke.go
*/

package chaninvoke

type Server struct {
	funcSet map[interface{}]interface{}
	CallChan chan * ArgPack
}

type Client struct {
	RetChan chan * RetPack
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
func (s *Server) Exec(arg *ArgPack) (err error) {
	result := arg.f.(func(...interface{}) interface{})(arg.args...)
	s.ret(arg, &RetPack{ret : result})
	return
}

func (s *Server) Call(id interface{}, args... interface{}){

}

func StartClient(l int, s *Server) *Client {
	c := new(Client)
	c.RetChan = make(chan *RetPack, l)
	c.s = s
	return c
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