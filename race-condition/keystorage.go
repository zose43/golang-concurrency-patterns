package main

import (
	"container/list"
	"sync"
)

// easily lru (least recently used) implementation

type Node struct {
	Data any
	Ptr  *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[any]*Node
	Capacity int
	mutex    sync.Mutex
}

func (r *LRUCache) Put(key any, value any) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if item, ok := r.Items[key]; !ok {
		if r.Capacity == r.Queue.Len() {
			back := r.Queue.Back()
			r.Queue.Remove(back)
			delete(r.Items, back.Value)
		}
		r.Items[key] = &Node{Data: value, Ptr: r.Queue.PushFront(key)}
	} else {
		item.Data = value
		r.Items[key] = item
		r.Queue.MoveToFront(item.Ptr)
	}
}

func (r *LRUCache) Get(key any) any {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if item, ok := r.Items[key]; ok {
		r.Queue.MoveToFront(item.Ptr)
		return item.Data
	}
	return nil
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Queue:    list.New(),
		Capacity: capacity,
		Items:    make(map[any]*Node),
	}
}
