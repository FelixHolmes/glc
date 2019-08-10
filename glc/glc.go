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
	// 业务函数
	BusinessFunc func()
	// 异常处理函数
	PanicFunc func()
	// 解码器
	Decoder func()
	// 编码器
	Encoder func()
}

// NewGLCServer 创建一个长连接服务
func NewGLCServer(conf ServerConfig) {

}
