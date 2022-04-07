package dstructs

import (
	"fmt"
	"strings"
)

type HashList[V comparable, K Node[K]] struct {
	mapList map[V]*List[K]

	Description string
	hashFunction func (*K) V
}

func NewHashList[V comparable, K Node[K]](description string, hf func (*K) V) *HashList[V,K] {

	hl := HashList[V,K]{}

	hl.mapList = make(map[V]*List[K])

	hl.hashFunction = hf

	return &hl
}

func (hl *HashList[V,K]) Add(n *K) {

	key := hl.hashFunction(n)

	list, ok := hl.mapList[key]
	if !ok {
		list = NewList[K](fmt.Sprintf("Key: %v", key))

		hl.mapList[key] = list
	}
	list.InsertHead(n)
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
