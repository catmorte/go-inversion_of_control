package context

import (
	"sync"
)

const defaultScope = ""

var currentContext Context
var lock sync.RWMutex

type dependencyRequest struct {
	Type   interface{}
	Waiter chan interface{}
	Scope  string
}

type Context interface {
	Ask(interfaceNil interface{}) chan interface{}
	Reg(interfaceNil interface{}, constructor func() interface{}, request ...*dependencyRequest)

	RegScoped(scope string, interfaceNil interface{}, constructor func() interface{}, request ...*dependencyRequest)
	AskScoped(scope string, interfaceNil interface{}) chan interface{}

	GetUnresolvedRequests() []*dependencyRequest
}

func Dep(t interface{}) *dependencyRequest {
	return &dependencyRequest{t, make(chan interface{}, 1), defaultScope}
}

func DepScoped(scope string, t interface{}) *dependencyRequest {
	return &dependencyRequest{t, make(chan interface{}, 1), scope}
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
