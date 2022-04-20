package dstructs

import (
	"fmt"
	"strings"
)

type HashList[V comparable, K Node[K]] struct {
	mapList map[V]*List[K]

	Description string

	// Function to produce hash value for keys
	hashFunction func (*K) V
	// Function to order items in List[K]
	orderFunction func (*K,*K) int
}

func NewHashList[V comparable, K Node[K]](description string, hf func (*K) V, of func (*K,*K) int) *HashList[V,K] {

	hl := HashList[V,K]{}

	hl.mapList = make(map[V]*List[K])

	hl.hashFunction = hf
	hl.orderFunction = of

	return &hl
}

func (hl *HashList[V,K]) Add(n *K) error {

	key := hl.hashFunction(n)

	list, ok := hl.mapList[key]
	if !ok {
		list = NewList(fmt.Sprintf("Key: %v", key), hl.orderFunction)

		hl.mapList[key] = list
	}

	if hl.orderFunction == nil {
		return list.InsertHead(n)
	} else {
		return list.InsertOrdered(n)
	}
}


func (hl *HashList[V,K]) Count() int {

	var count int
	for _, element := range hl.mapList {
		count += element.Len()
	}

	return count
}

func (hl *HashList[V,K]) GoString() string {

	var sb strings.Builder

	for _, element := range hl.mapList {
		sb.WriteString(fmt.Sprintf("%#v\n", element))
	}

	return sb.String()
}

func (hl *HashList[V,K]) String() string {

	var sb strings.Builder

	for _, element := range hl.mapList {
		sb.WriteString(fmt.Sprintf("%v\n", element))
	}

	return sb.String()
}

func (hl *HashList[V,K]) GetNodes() []*K {
	r := make([]*K, 0)

	for _, element := range hl.mapList {
		r = append(r, element.GetNodes()...)
	}

	return r
}
