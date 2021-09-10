package Map

import (
	"errors"
	"reflect"
	"sync"
)

var DiffKeyType = errors.New("[SafeMap]Unable diff key type")
var DiffValueType = errors.New("[SafeMap]Unable diff vale type")

//大颗粒锁
type SafeMap struct {
	m    map[interface{}]interface{}
	mu   sync.RWMutex
	kTyp reflect.Type //map key's type
	vTyp reflect.Type //map value's type
}

func New() *SafeMap {
	return &SafeMap{m: make(map[interface{}]interface{})}
}
func (s *SafeMap) Set(k, v interface{}) error {
	s.mu.Lock()
	if s.kTyp == nil {
		s.kTyp = reflect.TypeOf(k)
	}
	if s.vTyp == nil {
		s.vTyp = reflect.TypeOf(v)
	}
	if reflect.TypeOf(k) != s.kTyp {
		return DiffKeyType
	}
	if v != nil && reflect.TypeOf(v) != s.vTyp {
		return DiffValueType
	}
	s.m[k] = v
	s.mu.Unlock()
	return nil
}

func (s *SafeMap) Get(k interface{}) (interface{}, bool, error) {
	s.mu.RLock()
	if reflect.TypeOf(k) != s.kTyp {
		return nil, false, DiffKeyType
	}
	v, ok := s.m[k]
	s.mu.RUnlock()
	return v, ok, nil
}

func (s *SafeMap) Del(k interface{}) error {
	s.mu.Lock()
	if reflect.TypeOf(k) != s.kTyp {
		return DiffKeyType
	}
	delete(s.m, k)
	s.mu.Unlock()
	return nil
}

func (s *SafeMap) Amend(k, v interface{}) error {
	s.mu.Lock()
	if reflect.TypeOf(k) != s.kTyp {
		return DiffKeyType
	}
	if v != nil && reflect.TypeOf(v) != s.vTyp {
		return DiffValueType
	}
	s.m[k] = v
	s.mu.Unlock()
	return nil
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
