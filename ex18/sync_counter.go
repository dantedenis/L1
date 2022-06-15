package main

import (
	"sync"
)

type SyncCounter struct {
	sync.WaitGroup
	sync.RWMutex
	count int
}

func NewSyncCounter() *SyncCounter {
	return &SyncCounter{count:0}
}

func (s *SyncCounter) Inc() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *SyncCounter) GetCounter() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}