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

func Dep[T any]() *dependencyRequest {
	return &dependencyRequest{(*T)(nil), make(chan interface{}, 1), defaultScope}
}

func DepScoped[T any](scope string) *dependencyRequest {
	return &dependencyRequest{(*T)(nil), make(chan interface{}, 1), scope}
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

func Ask[T any](ctx Context) T {
	return (<-ctx.Ask((*T)(nil))).(T)
}

func Reg[T any](ctx Context, constructor func() interface{}, request ...*dependencyRequest) {
	ctx.Reg((*T)(nil), constructor, request...)
}

func RegScoped[T any](ctx Context, scope string, interfaceNil interface{}, constructor func() interface{}, request ...*dependencyRequest) {
	ctx.RegScoped(scope, (*T)(nil), constructor, request...)
}

func AskScoped[T any](ctx Context, scope string) T {
	return (<-ctx.AskScoped(scope, (*T)(nil))).(T)
}
