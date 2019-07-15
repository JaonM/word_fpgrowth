//FP tree
package core

//项头表元素
type HeadElem struct {
	word     string
	count    int
	treeNode *FPNode
	pattern  map[string]int // 条件模式基 {"word":"supportCount"}
}

type HeadElems [] HeadElem

func (elems HeadElems) Len() int {
	return len(elems)
}
func (elems HeadElems) Less(i, j int) bool {
	return elems[i].count < elems[j].count
}

func (elems HeadElems) Swap(i,j int) {
	elems[i],elems[j]=elems[j],elems[i]
}

type FPNode struct {
	word     string
	count    int
	child    *FPNode //Point to first child
	neighbor *FPNode // FPNode 邻接节点
	parent   *FPNode
	next     *FPNode //项头表指向节点
}

//FPTree 插入
func (p *FPNode) Insert(node *FPNode) {
	node.parent = p
	if p.child == nil {
		p.child = node
	} else {
		tmp := p.child
		for tmp.neighbor != nil {
			tmp = tmp.neighbor
		}
		tmp.neighbor = node
	}
}
