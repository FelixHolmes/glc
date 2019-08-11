package server

import (
	"github.com/FelixHolmes/glc/session"
	"net"
	"time"
)

type Server interface {
	Run() (*session.Session, error)
	Shutdown() error
	Use()
	ReStart() error
	Reload(conf Config) error
}
type Context struct {
}

// ServerConfig 服务通用配置
type Config struct {
	// 业务函数
	HandleFunc func(ctx Context)
	// 设置TCP的参数
	SetTcpFunc  func(conn *net.Conn)
	ConnLimit   int32
	ConnTimeOut time.Duration
}
