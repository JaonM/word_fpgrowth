package core

import (
	"testing"
)

func TestFPNode_Insert(t *testing.T) {
	root:=FPNode{}
	root.Insert(&FPNode{"apple",3,nil,nil,nil,nil})
	neighbor:=FPNode{"banana",4,nil,nil,nil,nil}
	root.Insert(&neighbor)
	if root.child.neighbor!=&neighbor {
		t.Error()
	}
}