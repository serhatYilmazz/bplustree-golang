package pkg

import (
	"fmt"
	"github.com/serhatYilmazz/bplustree/util"
	"math"
)

type BPTree interface {
	Add(key int, value int)
	PrintPreOrder(n *Node)
	//TODO add remove, update
}

type BpTree struct {
	Root   *Node
	Height int
}

const BranchingFactor = 5
const MaxKey = BranchingFactor - 1

var MinKey = int(math.Ceil(BranchingFactor/2)) - 1

type Node struct {
	Entry      *[BranchingFactor]Entry
	Children   *[BranchingFactor + 1]Node
	NodeLoad   int
	isInternal bool
}

type Entry struct {
	Key   int
	Value int
}

func (bt *BpTree) Add(key int, value int) {
	n := insert(key, value, bt.Height, bt.Height, bt.Root)
	if n == nil {
		return
	}

	newNode := &Node{
		Entry:      &[BranchingFactor]Entry{},
		Children:   &[BranchingFactor + 1]Node{},
		NodeLoad:   1,
		isInternal: true,
	}

	newNode.Entry[0] = Entry{
		Key:   n.Entry[0].Key,
		Value: 0,
	}

	if n.isInternal {
		arrangeInternal(n)
	}

	newNode.Children[0] = *bt.Root
	newNode.Children[1] = *n
	bt.Root = newNode
	bt.Height++
}

func insert(key int, value int, ht int, totalHt int, n *Node) *Node {
	j := 0
	e := Entry{
		Key:   key,
		Value: value,
	}
	// External Node
	if ht == 0 {
		for ; util.Less(n.Entry[j].Key, key) && j != n.NodeLoad; j++ {
		}
		for i := n.NodeLoad; i > j; i-- {
			n.Entry[i] = n.Entry[i-1]
		}
	} else { // Internal Node
		for ; util.Less(n.Entry[j].Key, key) && j != n.NodeLoad; j++ {
		}
		splittedNode := insert(key, value, ht-1, totalHt, &n.Children[j])
		if splittedNode == nil {
			return nil
		}
		e.Value = 0
		e.Key = splittedNode.Entry[0].Key
		for i := n.NodeLoad; i > j; i-- {
			n.Entry[i] = n.Entry[i-1]
			n.Children[i+1] = n.Children[i]
		}
		// If an internal node's key will be added to root node, it needs to be removed from internal
		if totalHt > 1 &&  ht == totalHt {
			arrangeInternal(splittedNode)
		}
		n.Children[j+1] = *splittedNode

	}

	n.Entry[j] = e
	n.NodeLoad++

	// Split
	if n.NodeLoad == BranchingFactor {
		return split(n)
	}
	return nil
}

func split(n *Node) *Node {
	median := BranchingFactor / 2
	node := &Node{
		Entry:      &[BranchingFactor]Entry{},
		Children:   &[BranchingFactor + 1]Node{},
		NodeLoad:   median + 1,
		isInternal: n.isInternal,
	}
	for i := 0; i < median+1; i++ {
		node.Entry[i] = n.Entry[i+median]
	}
	// There is overload, `branchingFactor` number of children exists, to reach them
	// For internal nodes
	for i := 0; i < n.NodeLoad-1; i++ {
		node.Children[i] = n.Children[i+median]
	}
	n.NodeLoad = median
	return node
}

func arrangeInternal(n *Node) {
	median := BranchingFactor / 2
	for i := 0; i < median+1; i++ {
		n.Entry[i] = n.Entry[i+1]
		n.Children[i] = n.Children[i+1]
	}
	n.NodeLoad = median
}

func (bt *BpTree) PrintPreOrder(n *Node, ht int) {
	if n.Entry == nil {
		return
	}
	for i := 0; i < n.NodeLoad; i++ {
		fmt.Printf("%d ", n.Entry[i].Key)
	}
	if n.isInternal {
		fmt.Printf("~ %d", ht)
	}
	fmt.Printf("\n")
	for i := 0; i < n.NodeLoad+1; i++ {
		bt.PrintPreOrder(&n.Children[i], ht-1)
	}

}
