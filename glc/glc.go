package glc

import "github.com/jueyoq/glc/session"

type Server interface {
	Run() (*session.Session, error)
	Shutdown() error
	ReStart() error
	Reload(conf ServerConfig) error
}

// ServerConfig 服务通用配置
type ServerConfig struct {
}

// NewGLCServer 创建一个长连接服务
func NewGLCServer(conf ServerConfig) {

}
