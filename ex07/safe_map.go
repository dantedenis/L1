package main

import "sync"

type SafeMap struct {
	sync.RWMutex
	data map[int]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[int]string),
	}
}

func (s *SafeMap) Add(key int, value string) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

func (s *SafeMap) GetValue(key int) string {
	s.RLock()
	defer s.RUnlock()
	return s.data[key]
}
