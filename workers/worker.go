package workers

type WorkerPool interface {
	Init(conf PoolConfig)
	Run()
	Close() error
	AddTask(func()) error
}

type PoolConfig struct {
}
