package context

import (
	"sync"
)

type memoryContext struct {
	storage  map[interface{}]interface{}
	requests map[interface{}][]chan interface{}
	lock     *sync.RWMutex
}

func (m *memoryContext) GetUnresolvedRequests() []*dependencyRequest {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var unresolvedRequests []*dependencyRequest
	for t, waiters := range m.requests {
		for _, waiter := range waiters {
			unresolvedRequests = append(unresolvedRequests, &dependencyRequest{t, waiter})
		}
	}
	return unresolvedRequests
}

func (m *memoryContext) Ask(t interface{}) chan interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()

	waiter := make(chan interface{})
	found := m.storage[t]
	if found != nil {

		go func() {
			waiter <- found
		}()
		return waiter
	}
	m.appendWaiter(t, waiter)
	return waiter
}

func (m *memoryContext) appendWaiter(t interface{}, waiter chan interface{}) {
	typ, ok := m.requests[t]
	if !ok {
		typ = []chan interface{}{}
		m.requests[t] = typ
	}
	m.requests[t] = append(typ, waiter)
}

func (m *memoryContext) Reg(t interface{}, constructor func() interface{}, requests ...*dependencyRequest) {
	m.lock.Lock()
	defer m.lock.Unlock()

	go func() {
		instance := constructor()
		m.lock.Lock()
		defer m.lock.Unlock()
		m.storage[t] = instance
		m.notify(t, instance)
	}()
	for _, r := range requests {
		found := m.storage[t]
		if found != nil {
			r.Waiter <- found
			continue
		}
		m.appendWaiter(r.Type, r.Waiter)
	}
}

func (m *memoryContext) notify(t interface{}, value interface{}) {

	if waiters, ok := m.requests[t]; ok {
		for _, w := range waiters {
			w <- value
		}
		delete(m.requests, t)
	}
}

func NewMemoryContext() Context {
	return &memoryContext{
		storage:  map[interface{}]interface{}{},
		requests: map[interface{}][]chan interface{}{},
		lock:     &sync.RWMutex{}}
}
