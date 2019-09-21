package context

import (
	"sync"
)

var currentContext Context
var lock sync.RWMutex

type dependencyRequest struct {
	Type   interface{}
	Waiter chan interface{}
}

type Context interface {
	Ask(interfaceNil interface{}) chan interface{}
	Reg(interfaceNil interface{}, constructor func() interface{}, request ...*dependencyRequest)
	GetUnresolvedRequests() []*dependencyRequest
}

func Dep(t interface{}) *dependencyRequest {
	return &dependencyRequest{t, make(chan interface{}, 1)}
}

func SetContext(context Context) {
	lock.Lock()
	defer lock.Unlock()
	currentContext = context
}

func GetContext() Context {
	lock.RLock()
	defer lock.RUnlock()
	return currentContext
}
