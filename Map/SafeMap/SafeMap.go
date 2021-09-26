package SafeMap

import (
	"sync"
)

//基于mutex RWLock的并发安全map
//想想还是panic()比较实在
//var DiffKeyType = errors.New("[SafeMap]Unable diff key type")
//var DiffValueType = errors.New("[SafeMap]Unable diff vale type")

//大颗粒锁
type SafeMap struct {
	m    map[interface{}]interface{}
	mu   sync.RWMutex
}

func New() *SafeMap {
	return &SafeMap{m: make(map[interface{}]interface{})}
}
func (s *SafeMap) Set(k, v interface{}) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func (s *SafeMap) Get(k interface{}) (interface{}, bool) {
	s.mu.RLock()
	v, ok := s.m[k]
	s.mu.RUnlock()
	return v, ok
}

func (s *SafeMap) Del(k interface{}) {
	s.mu.Lock()
	delete(s.m, k)
	s.mu.Unlock()
}

func (s *SafeMap) Amend(k, v interface{}) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func (s *SafeMap) Len() int {
	s.mu.RLock()
	l := len(s.m)
	s.mu.RUnlock()
	return l
}

type cb func(k, v interface{})

func (s *SafeMap) Range(callback cb) {
	s.mu.RLock()
	for k, v := range s.m {
		callback(k, v)
	}
	s.mu.RUnlock()
}

func (s *SafeMap) Copy() map[interface{}]interface{} {
	s.mu.RLock()
	m := s.m
	s.mu.RUnlock()
	return m
}

func (s *SafeMap) Elements() []interface{} {
	s.mu.RLock()
	var vs []interface{}
	for _, v := range s.m {
		vs = append(vs, v)
	}
	s.mu.RUnlock()
	return vs
}

func (s *SafeMap) Clean() {
	s.mu.Lock()
	s.m = map[interface{}]interface{}{}
	s.mu.Unlock()
}
