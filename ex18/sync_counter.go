package main

import (
	"sync"
)

// помимо самого счетчика втраиваем в структуру группы горутин и мьютексы
type SyncCounter struct {
	sync.WaitGroup
	sync.RWMutex
	count int
}

func NewSyncCounter() *SyncCounter {
	return &SyncCounter{count:0}
}

// конкурентная инкрементация за счет мьюта записи данных в структуру
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
