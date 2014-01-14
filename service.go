package runnerd

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
  "fmt"
)

const (
	CMD_START = iota
	CMD_STOP
	RESP_SUCCESS
	RESP_FAIL
)

type Request struct {
	Args []string
}

type Response struct {
	Status  int
	Message string
}

type ServiceConfig struct {
  Host string
  Port int
}

func (sc *ServiceConfig) HostPort() string {
  return fmt.Sprintf("%v:%v", sc.Host, sc.Port)
}

type Service struct {
  Config ServiceConfig
  Runner *Runner
}

type Client struct {
  Rpc *rpc.Client
}

func NewClient(port int) (*Client, error) {
  conf := ServiceConfig{"localhost", port}
	client, err := rpc.DialHTTP("tcp", conf.HostPort())
	if err != nil {
		log.Fatal("dialing:", err)
	}
  c := &Client{client}
  return c, nil
}

func (c *Client) Request(method string, args ...string) (*Response, error) {
  log.Println(method)
  r := &Request{args}
  w := &Response{}
  err := c.Rpc.Call(fmt.Sprintf("Service.%v", method), r, w)
  if err != nil {
    return nil, err
  }
  return w, nil
}

func (c *Client) Ping() (*Response, error) {
  w, err := c.Request("Ping")
  log.Println(w.Message)
  return w, err
}

func (c *Client) Info() (*Response, error) {
  w, err := c.Request("Info")
  log.Println(w.Message)
  return w, err
}

func (s *Service) Ping(r *Request, w *Response) error {
  w.Status = RESP_SUCCESS
  w.Message = "pong"
	return nil
}

func (s *Service) Info(r *Request, w *Response) error {
  if len(r.Args) == 0 {
    w.Status = RESP_SUCCESS
    w.Message = fmt.Sprintf("Info: %v", s.Runner.Ps)
  }
  return nil
}

func StartService(r *Runner) {
	s := new(Service)
  s.Runner = r
	rpc.Register(s)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

