package context

import (
	"sync"
)

const defaultScope = ""

var currentContext Context = NewMemoryContext()
var lock sync.RWMutex

type dependencyRequest struct {
	Type   any
	Waiter chan any
	Scope  string
}

type Context interface {
	Ask(interfaceNil any) chan interface{}
	Reg(interfaceNil any, constructor func() interface{}, request ...*dependencyRequest)

	RegScoped(scope string, interfaceNil any, constructor func() interface{}, request ...*dependencyRequest)
	AskScoped(scope string, interfaceNil any) chan interface{}

	GetUnresolvedRequests() []*dependencyRequest
}

func Dep[T any]() *dependencyRequest {
	return &dependencyRequest{(*T)(nil), make(chan any, 1), defaultScope}
}

func DepScoped[T any](scope string) *dependencyRequest {
	return &dependencyRequest{(*T)(nil), make(chan any, 1), scope}
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

func Ask[T any]() T {
	return (<-GetContext().Ask((*T)(nil))).(T)
}

func Reg[T any](constructor func() T, request ...*dependencyRequest) {
	GetContext().Reg((*T)(nil), typeToAnyFunc[T](constructor), request...)
}

func RegScoped[T any](scope string, constructor func() T, request ...*dependencyRequest) {
	GetContext().RegScoped(scope, (*T)(nil), typeToAnyFunc[T](constructor), request...)
}

func AskScoped[T any](scope string) T {
	return (<-GetContext().AskScoped(scope, (*T)(nil))).(T)
}

func ResolveDep[T any](dep *dependencyRequest) T {
	rawVal := <-dep.Waiter
	return (rawVal).(T)
}

func typeToAnyFunc[T any](f func() T) func() any {
	return func() any {
		return f()
	}
}
