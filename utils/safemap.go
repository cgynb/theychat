package utils

import "sync"

type SafeMap struct {
	m     map[any]any
	mutex sync.RWMutex
}

func (sm *SafeMap) Init() {
	sm.m = make(map[any]any)
}

func (sm *SafeMap) Set(key, value any) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.m[key] = value
}
func (sm *SafeMap) Get(key any) (value any, notEmpty bool) {
	sm.mutex.RLock()
	defer sm.mutex.RLocker()
	value, notEmpty = sm.m[key]
	return
}

func (sm *SafeMap) Del(key any) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	delete(sm.m, key)
}
