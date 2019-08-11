package tcpserver

import (
	"github.com/FelixHolmes/glc/server"
	"github.com/FelixHolmes/glc/session"
)

func NewTcpServer(conf server.Config) server.Server {
	return &TcpServer{}
}

type TcpServer struct {
}

func (s *TcpServer) Run() (*session.Session, error) {
	return nil, nil
}

func (s *TcpServer) Shutdown() error {
	return nil
}
func (s *TcpServer) Use() {

}
func (s *TcpServer) ReStart() error {
	return nil
}
func (s *TcpServer) Reload(conf server.Config) error {
	return nil
}
