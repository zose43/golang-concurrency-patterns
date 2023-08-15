package main

import (
	"container/list"
	"sync"
)

// easily lru (least recently used) implementation

type Node struct {
	Data string
	Ptr  *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[int]*Node
	Capacity int
	mutex    sync.Mutex
}

func (r *LRUCache) Put(key int, value string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if item, ok := r.Items[key]; !ok {
		if r.Capacity == r.Queue.Len() {
			back := r.Queue.Back()
			r.Queue.Remove(back)
			delete(r.Items, back.Value.(int))
		}
		r.Items[key] = &Node{Data: value, Ptr: r.Queue.PushFront(key)}
	} else {
		item.Data = value
		r.Items[key] = item
		r.Queue.MoveToFront(item.Ptr)
	}
}

func (r *LRUCache) Get(key int) string {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if item, ok := r.Items[key]; ok {
		r.Queue.MoveToFront(item.Ptr)
		return item.Data
	}
	return ""
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Queue:    list.New(),
		Capacity: capacity,
		Items:    make(map[int]*Node),
	}
}
