package hashid

import (
	"sync"
)

type HashID interface {
	Get(string) int64
	Exist(string) bool
	Len() int64
}

type HashIDImpl struct {
	data map[string]int64
	id   int64
	mu   sync.RWMutex
}

func NewHashID() HashID {
	return &HashIDImpl{
		data: map[string]int64{},
		id:   0,
	}
}

func (h *HashIDImpl) Get(key string) int64 {
	h.mu.Lock()
	defer h.mu.Unlock()
	_, ok := h.data[key]
	if !ok {
		h.data[key] = h.id
		h.id++
	}
	return h.data[key]
}

func (h *HashIDImpl) Exist(key string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	_, ok := h.data[key]
	return ok
}

func (h *HashIDImpl) Len() int64 {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.id
}
