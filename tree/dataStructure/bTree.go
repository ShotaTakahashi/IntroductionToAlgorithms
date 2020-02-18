package dataStructure

import (
	"fmt"
	"sort"
)

type BTreeNode struct {
	n        int
	keys     []int
	leaf     bool
	children map[int]*BTreeNode
}

type BTree struct {
	Root *BTreeNode
}

func (node *BTreeNode) NumberOfNodes() int {
	return node.n
}

func (node *BTreeNode) GetKeyList() []int {
	return node.keys
}

func (node *BTreeNode) IsLeaf() bool {
	return node.leaf
}

func (node *BTreeNode) AddChildren(children *BTreeNode, index int) {
	node.children[index] = children
}

func (node *BTreeNode) GetChildren() map[int]*BTreeNode {
	return node.children
}

func NewBTreeNode(keys []int, leaf bool) *BTreeNode {
	node := &BTreeNode{}
	node.n = len(keys)
	sort.Ints(keys)
	node.keys = keys
	node.leaf = leaf
	node.children = make(map[int]*BTreeNode)
	return node
}

func (tree *BTree) Search(key int) (node *BTreeNode, index int) {
	return bTreeSearch(tree.Root, key)
}

func bTreeSearch(root *BTreeNode, key int) (node *BTreeNode, index int) {
	index = 0
	for index <= root.n && key > root.keys[index] {
		index++
	}
	if index <= root.n && key == root.keys[index] {
		node = root
		return
	} else if root.leaf {
		return nil, -1
	} else {
		return bTreeSearch(root.children[index], key)
	}
}

func (tree *BTree) Insert(key int) {
	r := tree.Root
	t := r.n/2 + 1
	if r.n == 2*t-1 {
		s := &BTreeNode{}
		tree.Root = s
		s.leaf = false
		s.n = 0
		s.children = make(map[int]*BTreeNode)
		s.children[0] = r
		BTreeSplitChild(s, 0)
		fmt.Println(s.children[1])
		bTreeInsertNonfull(s, key)
	} else {
		bTreeInsertNonfull(r, key)
	}
}

func bTreeInsertNonfull(node *BTreeNode, key int) {
	t := node.n/2 + 1
	i := node.n - 1
	if node.leaf {
		for i >= 0 && key < node.keys[i] {
			//node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys = append(node.keys, key)
		sort.Ints(node.keys)
		node.n++
	} else {
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++
		if node.children[i] == nil {
			node.children[i] = NewBTreeNode([]int{key}, true)
		} else if node.children[i].n == 2*t-1 {
			BTreeSplitChild(node, i)
			if key > node.keys[i] {
				i++
			}
			bTreeInsertNonfull(node.children[i], key)
		} else {
			bTreeInsertNonfull(node.children[i], key)
		}
	}
}

func NewBTree() (bTree *BTree) {
	root := NewBTreeNode([]int{}, true)
	bTree = &BTree{}
	bTree.Root = root
	return
}

func NewBTreeWithNode(node *BTreeNode) (bTree *BTree) {
	bTree = &BTree{}
	bTree.Root = node
	return
}

func BTreeSplitChild(node *BTreeNode, index int) {
	y := node.children[index]
	z := &BTreeNode{}
	z.leaf = y.leaf
	t := y.n/2 + 1
	z.n = t - 1
	z.keys = make([]int, z.n)
	z.children = make(map[int]*BTreeNode)
	for j := 0; j < t-1; j++ {
		z.keys[j] = y.keys[j+t]
	}
	if !y.leaf {
		for j := 0; j < t; j++ {
			z.children[j] = y.children[j+t]
		}
	}
	y.n = t - 1
	for j := node.n + 1; j > index+1; j-- {
		node.children[j+1] = node.children[j]
	}
	node.children[index+1] = z
	nkeys := make([]int, node.n+1)
	for j := node.n - 1; j > index-1; j-- {
		nkeys[j+1] = node.keys[j]
	}
	for j := 0; j < index; j++ {
		nkeys[j] = node.keys[j]
	}
	nkeys[index] = y.keys[t-1]
	node.n++
	node.keys = nkeys
	y.keys = y.keys[:y.n]
}
