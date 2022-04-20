package dstructs

import (
	"fmt"
	"strings"
)

type linkNode[T Node[T]] struct {
	Data *T
	nextNode *linkNode[T]
}

func (ln *linkNode[T]) next() *linkNode[T] {
	return ln.nextNode
}

func (ln *linkNode[T]) setNext(next *linkNode[T]) {
	ln.nextNode = next
}

func (ln *linkNode[T]) getData() *T {
	return ln.Data
}

func newLinkNode[T Node[T]](data *T) *linkNode[T] {
	return &linkNode[T]{
		Data: data,
		nextNode: nil,
	}
}

type List[T Node[T]] struct {
	Head linkNode[T]
	Tail linkNode[T]

	Description string

	OrderFunc func(*T,*T) int
}

func NewList[T Node[T]](description string,fn func(*T,*T) int) *List[T] {
	
	return &List[T]{
		Description: description,
		OrderFunc: fn,
	}
}

func (l *List[T]) Len() int {
	length := 0

	for node := &l.Head; node.next() != nil; node = node.next() {
		length++
	}

	return length
}

func (l *List[T]) InsertHead(data *T) error {
	nn := newLinkNode(data)

	nn.setNext(l.Head.next())
	if l.Head.next() == nil {
		l.Tail.setNext(nn)
	}
	l.Head.setNext(nn)

	return nil
}

func (l *List[T]) InsertTail(data *T) {
	nn := newLinkNode(data)

	nn.setNext(l.Tail.next())
	if l.Tail.next() == nil {
		l.Head.setNext(nn)
	}
	l.Tail.setNext(nn)
}

func (l *List[T]) InsertOrdered(data *T) error {
	if l.OrderFunc == nil {
		return fmt.Errorf("not order function present")
	}

	nn := newLinkNode(data)

	var node *linkNode[T]

	for node = &l.Head; node != nil; node = node.next() {
		nextNode := node.next()
		if nextNode == nil {
			node.setNext(nn)
			l.Tail.setNext(nn)
			break
		}
		nextNodeData := nextNode.getData()
		if l.OrderFunc(data,nextNodeData) <= 0 {
			nn.setNext(nextNode)
			node.setNext(nn)
			break
		}
	}

	return nil
}

func (l *List[T]) Filter(fn func(*T) int) []*T {
	r := make([]*T, 0)

	for e := l.Head.next(); e != nil; e = e.next() {
		data := e.getData()
		if fn(data) == 0 {
			r = append(r, data)
		}
	}

	return r
}

func (l *List[T]) Get(n *T) *T {
	for e := l.Head.next(); e != nil; e = e.next() {
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
	for e := l.Head.next(); e != nil; e = e.next() {
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
	for e := l.Head.next(); e != nil; e = e.next() {
		sb.WriteString(fmt.Sprintf("%v ", e.getData()))
	}
	sb.WriteString("]")

	return sb.String()
}

func (l *List[T]) GetNodes() []*T {
	r := make([]*T, 0)

	for e := l.Head.next(); e != nil; e = e.next() {
		data := e.getData()
		r = append(r, data)
	}

	return r
}
