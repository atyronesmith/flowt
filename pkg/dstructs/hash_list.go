package dstructs

type HashList[V comparable, K Node[V,K]] struct {
	mapList map[int]*List[V,K]

}

func NewHashList[V comparable, K Node[V,K]]() *HashList[V,K] {

	hl := HashList[V,K]{}

	hl.mapList = make(map[int]*List[V,K])

	return &hl
}

func (hl *HashList[V,K]) Add(n *K) {

	list, ok := hl.mapList[0]
	if !ok {
		list = NewList[V,K]()

		hl.mapList[0] = list
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
