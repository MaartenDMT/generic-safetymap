package safemap

import "sync"

type SafeMap[K comparable , V any] struct{
	mu sync.RWMutex
	data map[K]V

}

func New[K comparable, V any] *SafeMap[K,V]{
	return &SafeMap[K,V]{
        data: make(map[K]V),
    }
}

func (m *SafeMap[K,V]) Insert(key K, value V){
	m.mu.Lock()
    defer m.mu.Unlock()
    m.data[key] = value
}

func (m *SafeMap[K,V]) Delete(key K){
	m.mu.Lock()
    defer m.mu.Unlock()
    delete(m.data, key)
}

func (m *SafeMap[K,V]) Get(key K){
	m.mu.RLock()
    defer m.mu.RUnlock()
    return m.data[key]
}