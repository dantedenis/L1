package main

import "sync"

type SafeMap struct {
	// Встраиваем структуру RWMutex, более оптимальный процесс записи и чтения
	sync.RWMutex
	data map[int]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[int]string),
	}
}

func (s *SafeMap) Add(key int, value string) {
	//блокирование мьютекса на каждую запись в мапу
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

func (s *SafeMap) GetValue(key int) string {
	//блокирование мьютекса на чтение  и последующий доступ в мапу
	s.RLock()
	defer s.RUnlock()
	return s.data[key]
}
