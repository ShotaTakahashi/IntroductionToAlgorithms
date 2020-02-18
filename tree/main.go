package main

import (
	"fmt"

	ds "github.com/ShotaTakahashi/IntroductionToAlgorithms/tree/dataStructure"
)

func main() {
	n0 := ds.Node{Key: 2, Left: nil, Right: nil}
	n1 := ds.Node{Key: 5, Left: nil, Right: nil}
	n2 := ds.Node{Key: 5, Left: &n0, Right: &n1}
	n3 := ds.Node{Key: 8, Left: nil, Right: nil}
	n4 := ds.Node{Key: 7, Left: nil, Right: &n3}
	root := ds.Node{Key: 6, Left: &n2, Right: &n4}
	fmt.Println(ds.IsBinarySearchTree(root))
	fmt.Println(ds.TreeSearch(root, 3))

	fmt.Println("BTree")
	bRoot := ds.NewBTreeNode([]int{4, 10, 25, 42}, false)
	bTree := ds.NewBTreeWithNode(bRoot)
	bChild0 := ds.NewBTreeNode([]int{-1, 1, 3}, true)
	bTree.Root.AddChildren(bChild0, 0)
	bChild1 := ds.NewBTreeNode([]int{6, 8, 9}, true)
	bTree.Root.AddChildren(bChild1, 1)
	bChild2 := ds.NewBTreeNode([]int{14}, true)
	bTree.Root.AddChildren(bChild2, 2)
	bChild3 := ds.NewBTreeNode([]int{28, 32, 33, 36, 39}, true)
	bTree.Root.AddChildren(bChild3, 3)

	fmt.Println(bTree.Root.GetKeyList())
	node, index := bTree.Search(4)
	fmt.Println(*node, index)
	fmt.Println(ds.NewBTree())

	bTree.Insert(45)
	fmt.Println(bTree.Root.GetKeyList())
	fmt.Println(bTree.Root.GetChildren()[0],
		bTree.Root.GetChildren()[1],
		bTree.Root.GetChildren()[2],
		bTree.Root.GetChildren()[3],
		bTree.Root.GetChildren()[4])
}
