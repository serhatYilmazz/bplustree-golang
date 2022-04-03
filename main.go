package main

import (
	"github.com/serhatYilmazz/bplustree/pkg"
)

func main() {
	bt := pkg.BpTree{
		Root:   &pkg.Node{
			Entry:    &[pkg.BranchingFactor]pkg.Entry{},
			Children: &[pkg.BranchingFactor + 1]pkg.Node{},
			NodeLoad: 0,
		},
		Height: 0,
	}

	bt.Add(10, 1)
	bt.Add(20, 2)
	bt.Add(30, 2)
	bt.Add(40, 2)
	bt.Add(50, 2)
	bt.Add(60, 1)
	bt.Add(70, 2)
	bt.Add(41, 2)
	bt.Add(42, 2)
	bt.Add(43, 2)
	bt.Add(31, 2)
	bt.Add(32, 2)
	bt.Add(33, 2)
	bt.Add(11, 1)
	bt.Add(12, 2)
	bt.Add(13, 2)

	bt.Add(80, 2)
	bt.Add(81, 2)
	bt.Add(44, 2)
	bt.Add(45, 2)
	bt.Add(34, 2)
	bt.Add(35, 2)

	bt.Add(14, 2)
	bt.Add(15, 2)
	bt.Add(16, 2)
	bt.Add(17, 2)
	bt.Add(18, 2)
	bt.Add(19, 2)

	bt.Add(46, 2)
	bt.Add(47, 2)
	bt.Add(48, 2)
	bt.Add(49, 2)
	bt.Add(51, 2)
	bt.Add(52, 2)
	bt.Add(53, 2)

	bt.Add(54, 2)
	bt.Add(55, 2)
	bt.Add(56, 2)
	bt.Add(57, 2)
	bt.Add(58, 2)
	bt.Add(59, 2)


	bt.PrintPreOrder(bt.Root, bt.Height)
}
