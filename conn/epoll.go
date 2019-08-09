package conn

type Conn struct {
}
type LakeConfig struct {
}
type Lake interface {
	Init(conf LakeConfig) error
	Run() error
}
