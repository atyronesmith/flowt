package dstructs

type Node[T any] interface {
	Compare(*T) int
}
