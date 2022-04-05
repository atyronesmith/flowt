package dstructs

import (
	"fmt"
)

type linkNode[V comparable, T Node[V,T]] struct {
	Data *T
	Next *linkNode[V,T]
}

func (ln *linkNode[V,T]) next() *linkNode[V,T] {
	return ln.Next
}

func (ln *linkNode[V,T]) setNext(next *linkNode[V,T]) {
	ln.Next = next
}

func (ln *linkNode[V,T]) getData() *T {
	return ln.Data
}

func newLinkNode[V comparable, T Node[V,T]](data *T) *linkNode[V,T] {
	return &linkNode[V,T]{
		Data: data,
		Next: nil,
	}
}

type List[V comparable, T Node[V,T]] struct {
	Head *linkNode[V,T]
	Tail *linkNode[V,T]
}

func NewList[V comparable, T Node[V,T]]() *List[V,T] {
	return &List[V,T]{
		Head: nil,
		Tail: nil,
	}
}

func (l *List[V,T]) Len() int {
	length := 0
	for e := l.Head; e != nil; e = e.next() {
		length++
	}

	return length
}

func (l *List[V,T]) InsertHead(data *T) {
	nn := newLinkNode[V](data)
	
	nn.setNext(l.Head)
	if l.Head == nil {
		l.Tail = nn
	}
	l.Head = nn
}

func (l *List[V,T]) InsertTail(data *T) {
	nn := newLinkNode[V](data)

	nn.setNext(l.Tail)
	if l.Tail == nil {
		l.Head = nn
	}
	l.Tail = nn
}

func (l *List[V,T]) Filter(fn func(*T) int) []*T {
	r := make([]*T, 0)

	for e := l.Head; e != nil; e = e.next() {
		data := e.getData()
		if fn(data) == 0 {
			r = append(r, data)
		}
	}

	return r
}

func (l *List[V,T]) Get(n *T) *T {
	for e := l.Head; e != nil; e = e.next() {
		data := e.getData()			
		if (*data).Compare(n) == 0 {
			return data
		}
	}
	return nil
}

func (l *List[V,T]) String() string {
	r := ""

	for e := l.Head; e != nil; e = e.next() {
		r += fmt.Sprintf("%v", e.getData())
	}

	return r
}
