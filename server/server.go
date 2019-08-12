package server

import (
	"github.com/FelixHolmes/glc/session"
	"net"
	"time"
)

type HandleFunc func(ctx *Context)
type Server interface {
	Run() (*session.Session, error)
	Shutdown() error
	Use(h HandleFunc)
	ReStart() error
	Reload(conf Config) error
}
type Context struct {
	index int
}

func (c *Context) Next() {
	c.index++
	//TODO:拿不到上下文
}

// ServerConfig 服务通用配置
type Config struct {
	// 业务函数
	HandleFunc func(ctx *Context)
	// 设置TCP的参数
	SetTcpFunc  func(conn *net.Conn)
	ConnLimit   int32
	ConnTimeOut time.Duration
}
