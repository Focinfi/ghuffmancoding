package ghuffmancoding

import (
	"sort"
)

// Node is a unit of the huffman tree
type Node struct {
	Value      string
	Weight     int
	LeftChild  *Node
	RightChild *Node
}

// traverse visit the whole substree from this node
func (n Node) traverse(code string, visit func(string, string)) {
	if leftNode := n.LeftChild; leftNode != nil {
		leftNode.traverse(code+"0", visit)
	} else {
		visit(n.Value, code)
		return
	}
	n.RightChild.traverse(code+"1", visit)
}

type Nodes []Node

type Tree struct {
	Root *Node
}

// encode traverse from the root of the tree and put the encoding result into a map
func (tree Tree) encode() map[string]string {
	var initialCode string
	encodeMap := make(map[string]string)
	tree.Root.traverse(initialCode, func(value string, code string) {
		encodeMap[value] = code
	})
	return encodeMap
}

// Len implements Len() int in sort.Interface
func (h Nodes) Len() int {
	return len(h)
}

// Less implements Less(i, j int) bool in sort.Interface
func (h Nodes) Less(i, j int) bool {
	return h[i].Weight > h[j].Weight
}

// Swap implements Swap(i, j int) int in sort.Interface
func (h Nodes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Encode encode a str into a map[string]string
// Example
//   result := huffmancoding.Encode("111223") // result: map[3:00 2:01 1:1]
func Encode(str string) map[string]string {
	priorityMap := makePriorityMap(str)
	stortedNodes := makeSortedNodes(priorityMap)
	hfmTree := makeFuffManTree(stortedNodes)
	return hfmTree.encode()
}

// makePriorityMap make a map[string]int
// key is the distinct character in string, value is the key's times of appration
func makePriorityMap(str string) map[string]int {
	matchMap := make(map[string]int)
	for _, chr := range str {
		matchMap[string(chr)] += 1
	}
	return matchMap
}

// makeSortedNodes make a []Node ordered by ascending Weight
func makeSortedNodes(priorityMap map[string]int) []Node {
	hfmNodes := make(Nodes, len(priorityMap))
	i := 0
	for value, weight := range priorityMap {
		hfmNodes[i] = Node{Value: value, Weight: weight}
		i++
	}
	sort.Sort(sort.Reverse(hfmNodes))
	return hfmNodes
}

// makeFuffManTree make a huffman tree using the sorted node array
func makeFuffManTree(nodes Nodes) *Tree {
	if len(nodes) < 2 {
		panic("Must contain 2 or more emlments")
	}
	hfmTree := &Tree{&Node{"", nodes[0].Weight + nodes[1].Weight, &nodes[0], &nodes[1]}}
	for i := 2; i < len(nodes); {
		if nodes[i].Weight == 0 {
			i++
			continue
		}
		oldRoot := hfmTree.Root
		if i+1 < len(nodes) && hfmTree.Root.Weight > nodes[i+1].Weight {
			newNode := Node{"", nodes[i].Weight + nodes[i+1].Weight, &nodes[i], &nodes[i+1]}
			hfmTree.Root = &Node{"", newNode.Weight + oldRoot.Weight, oldRoot, &newNode}
			i += 2
		} else {
			hfmTree.Root = &Node{"", nodes[i].Weight + oldRoot.Weight, oldRoot, &nodes[i]}
			i++
		}
	}
	return hfmTree
}
