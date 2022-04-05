package dstructs

// Node interface defines an interface for nodes in the linked list
// so we can iterate them easier.
type Node[V comparable, T any] interface {
	Compare(*T) int
	Key() V
}
