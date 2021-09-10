package Slice

import (
	"errors"
	"reflect"
	"sync"
)

var ErrOutOfBounds = errors.New("[SafeSlice]Index out of bounds")
var ErrSliceFull = errors.New("[SafeSlice]Cap no empty")
var ErrDiffType = errors.New("[SafeSlice]Unable diff type")

type SafeSlice struct {
	slice []interface{}
	count int
	cap   int
	mu    sync.RWMutex
	typ   reflect.Type //Elements in slice must be same type
}

//variable length array only effective when cap > 0
func New(cap ...int) *SafeSlice {
	if len(cap) != 0 {
		c := cap[0]
		if c > 0 {
			return &SafeSlice{
				slice: make([]interface{}, c),
				count: 0,
			}
		}
	}
	return &SafeSlice{
		slice: []interface{}{},
		count: 0,
	}
}

func (s *SafeSlice) Append(vs ...interface{}) error {
	s.mu.Lock()
	if s.cap != 0 && s.cap >= s.count {
		return ErrSliceFull
	}
	for _, v := range vs {
		if v != nil {
			if s.typ == nil {
				s.typ = reflect.TypeOf(v)
			} else {
				if s.typ != reflect.TypeOf(v) {
					return ErrDiffType
				}
			}
		}
	}
	s.slice = append(s.slice, vs...)
	s.count++
	s.mu.Unlock()
	return nil
}

func (s *SafeSlice) Amend(i int, v interface{}) error {
	s.mu.Lock()
	if i >= len(s.slice) || i < 0 {
		return ErrOutOfBounds
	}
	if v != nil {
		if s.typ != reflect.TypeOf(v) {
			return ErrDiffType

		}
	}
	s.slice[i] = v
	s.mu.Unlock()
	return nil
}

func (s *SafeSlice) Index(i int) (interface{}, error) {
	s.mu.RLock()
	if i >= len(s.slice) || i < 0 {
		return nil, ErrOutOfBounds
	}
	v := s.slice[i]
	s.mu.RUnlock()
	return v, nil
}

func (s *SafeSlice) Del(i int) error {
	s.mu.Lock()
	if i >= s.count || i < 0 {
		return ErrOutOfBounds
	}
	if i == 0 {
		s.slice = s.slice[1:]
	} else if i == s.count-1 {
		s.slice = s.slice[:i-1]
	} else {
		s.slice = append(s.slice[:i], s.slice[i+1:]...)
	}
	s.mu.Unlock()
	return nil
}

func (s *SafeSlice) Len() int {
	s.mu.RLock()
	l := s.count
	s.mu.RUnlock()
	return l
}

type cb func(i int, v interface{})

func (s *SafeSlice) Range(callback cb) {
	s.mu.RLock()
	for i, v := range s.slice {
		callback(i, v)
	}
	s.mu.RUnlock()
}

func (s *SafeSlice) All() []interface{} {
	s.mu.RLock()
	cp := s.slice
	s.mu.RUnlock()
	return cp
}
