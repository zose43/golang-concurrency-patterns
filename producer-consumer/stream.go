package main

import "time"

type Stream struct {
	airport  string
	timezone string
	city     string
	date     time.Time
	next     *Stream
}

type LinkedList struct {
	head   *Stream
	length int
}

func (l *LinkedList) reverse() {
	var cur, prev, next *Stream
	cur = l.head
	prev = nil
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head = prev
}

func (l *LinkedList) removeFromHead() {
	l.head = l.head.next
	l.length--
}

func (l *LinkedList) remove(n int) {
	switch n {
	case 0:
		l.removeFromHead()
	case l.length - 1:
		l.removeFromBack()
	default:
		cur := l.head
		var next *Stream
		for i := 0; i < n; i++ {
			next = cur.next
		}
		cur.next = next.next
	}
	l.length--
}

func (l *LinkedList) removeFromBack() {
	cur := l.head
	for i := 0; i < l.length-2; i++ {
		cur = cur.next
	}
	cur.next = nil
	l.length--
}

func (l *LinkedList) insert(n int, stream *Stream) {
	switch n {
	case 0:
		l.addToHead(stream)
	case l.length - 1:
		l.addToBack(stream)
	default:
		cur := l.head
		for i := 0; i < n-1; i++ {
			cur = cur.next
		}
		stream.next = cur.next
		cur.next = stream
	}
	l.length++
}

func (l *LinkedList) addToHead(stream *Stream) {
	if l.head == nil {
		l.head = stream
	} else {
		oldhead := l.head
		l.head = stream
		stream.next = oldhead
	}
	l.length++
}

func (l *LinkedList) addToBack(stream *Stream) {
	if l.head == nil {
		l.head = stream
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = stream
	}
	l.length++
}
