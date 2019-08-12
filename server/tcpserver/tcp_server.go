package tcpserver

import (
	"github.com/FelixHolmes/glc/server"
	"github.com/FelixHolmes/glc/session"
)

func NewTcpServer(conf server.Config) server.Server {
	return &TcpServer{}
}

type TcpServer struct {
	handlers []server.HandleFunc
	task     server.HandleFunc
}

func (s *TcpServer) Run() (*session.Session, error) {
	s.mergeFunc()
	return nil, nil
}

func (s *TcpServer) Shutdown() error {
	return nil
}
func (s *TcpServer) Use(h server.HandleFunc) {
	s.handlers = append(s.handlers, h)
}
func (s *TcpServer) ReStart() error {
	return nil
}
func (s *TcpServer) Reload(conf server.Config) error {
	return nil
}

func (s *TcpServer) mergeFunc() {
	s.handlers = append(s.handlers, s.task)
}
