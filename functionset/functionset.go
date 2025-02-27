package functionset

import "sync"

type FunctionSet struct {
	fn []func(...interface{})
	*syncFunc
}

type syncFunc struct {
	wg      sync.WaitGroup
	cond    sync.Cond
	mp      sync.Map
	mutex   sync.Mutex
	once    sync.Once
	pool    sync.Pool
	rwmutex sync.RWMutex
	locker  sync.Locker
}

func NewFunctionSet(fns ...func(...interface{})) *FunctionSet {
	return &FunctionSet{
		fn: fns,
	}
}
