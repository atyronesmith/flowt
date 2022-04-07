package dstructs

import (
	"fmt"
	"strings"
)

type linkNode[T Node[T]] struct {
	Data *T
	Next *linkNode[T]
}

func (ln *linkNode[T]) next() *linkNode[T] {
	return ln.Next
}

func (ln *linkNode[T]) setNext(next *linkNode[T]) {
	ln.Next = next
}

func (ln *linkNode[T]) getData() *T {
	return ln.Data
}

func newLinkNode[T Node[T]](data *T) *linkNode[T] {
	return &linkNode[T]{
		Data: data,
		Next: nil,
	}
}

type List[T Node[T]] struct {
	Head *linkNode[T]
	Tail *linkNode[T]

	Description string
}

func NewList[T Node[T]](description string) *List[T] {
	return &List[T]{
		Head:        nil,
		Tail:        nil,
		Description: description,
	}
}

func (l *List[T]) Len() int {
	length := 0
	for e := l.Head; e != nil; e = e.next() {
		length++
	}

	return length
}

func (l *List[T]) InsertHead(data *T) {
	nn := newLinkNode(data)

	nn.setNext(l.Head)
	if l.Head == nil {
		l.Tail = nn
	}
	l.Head = nn
}

func (l *List[T]) InsertTail(data *T) {
	nn := newLinkNode(data)

	nn.setNext(l.Tail)
	if l.Tail == nil {
		l.Head = nn
	}
	l.Tail = nn
}

func (l *List[T]) Filter(fn func(*T) int) []*T {
	r := make([]*T, 0)

	for e := l.Head; e != nil; e = e.next() {
		data := e.getData()
		if fn(data) == 0 {
			r = append(r, data)
		}
	}

	return r
}

func (l *List[T]) Get(n *T) *T {
	for e := l.Head; e != nil; e = e.next() {
		data := e.getData()
		if (*data).Compare(n) == 0 {
			return data
		}
	}
	return nil
}

func (l *List[T]) GoString() string {
	var sb strings.Builder

	if len(l.Description) > 0 {
		sb.WriteString(fmt.Sprintf("[ %s ]\n", l.Description))
	}
	for e := l.Head; e != nil; e = e.next() {
		sb.WriteString(fmt.Sprintf("%#v", e.getData()))
		sb.WriteString("\n")
	}

	return sb.String()
}

func (l *List[T]) String() string {
	var sb strings.Builder

	if len(l.Description) > 0 {
		sb.WriteString(fmt.Sprintf("[ %s, ", l.Description))
	}
	for e := l.Head; e != nil; e = e.next() {
		sb.WriteString(fmt.Sprintf("%v ", e.getData()))
	}
	sb.WriteString("]")

	return sb.String()
}

func (l *List[T]) GetNodes() []*T {
	r := make([]*T, 0)

	for e := l.Head; e != nil; e = e.next() {
		data := e.getData()
		r = append(r, data)
	}

	return r
}
