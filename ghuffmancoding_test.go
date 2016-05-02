package ghuffmancoding

import (
	"testing"
)

func TestNode(t *testing.T) {
	Node := Node{Weight: 1}
	if Node.Weight != 1 {
		t.Error("Can not instructed")
	}
}

func TestTree(t *testing.T) {
	RootRode := Node{Weight: 1}
	Tree := Tree{&RootRode}
	if Tree.Root == nil {
		t.Error("Can not instrcuted")
	}
}

func TestMakePriorityMap(t *testing.T) {
	str := "1123145512"
	priorityMap := makePriorityMap(str)
	if priorityMap['1'] != 4 {
		t.Error("Can not make a right map, priorityMap[1] =", priorityMap['1'])
	}
}

func TestMakeSortedNodes(t *testing.T) {
	str := "112"
	priorityMap := makePriorityMap(str)
	stortedNodes := makeSortedNodes(priorityMap)
	if stortedNodes[0].Value != '2' || stortedNodes[1].Value != '1' {
		t.Error("Can not sort Map, sortedNodes[0] is:", stortedNodes[0])
	}
}

func TestMakeFuffManTree(t *testing.T) {
	str := "111223"
	priorityMap := makePriorityMap(str)
	stortedNodes := makeSortedNodes(priorityMap)
	hfmTree := makeFuffManTree(stortedNodes)
	if hfmTree.Root.Weight != 6 {
		t.Error("Can not make a hfmTree, root is:", hfmTree.Root.LeftChild, hfmTree.Root.RightChild)
	}
}

func TestTraverse(t *testing.T) {
	str := "111223"
	encoding := Encode(str)
	if encoding['1'] != "1" || encoding['2'] != "01" || encoding['3'] != "00" {
		t.Error("Can not traverse in pre order, first element is:", encoding)
	}
}
