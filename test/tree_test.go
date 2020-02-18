package test

import (
	"testing"

	ds "github.com/ShotaTakahashi/IntroductionToAlgorithms/tree/dataStructure"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	n0 := ds.Node{Key: 2, Left: nil, Right: nil}
	n1 := ds.Node{Key: 5, Left: nil, Right: nil}
	n2 := ds.Node{Key: 5, Left: &n0, Right: &n1}
	n3 := ds.Node{Key: 8, Left: nil, Right: nil}
	n4 := ds.Node{Key: 7, Left: nil, Right: &n3}
	root := ds.Node{Key: 6, Left: &n2, Right: &n4}

	assert.Equal(t, true, ds.IsBinarySearchTree(root))

	n3 = ds.Node{Key: 4, Left: nil, Right: nil}

	assert.Equal(t, false, ds.IsBinarySearchTree(root))

	n3 = ds.Node{Key: 8, Left: nil, Right: nil}

	_, isInTree := ds.TreeSearch(root, 3)
	assert.Equal(t, false, isInTree)

	node, isInTree := ds.TreeSearch(root, 5)
	assert.Equal(t, &n2, node)
	assert.Equal(t, true, isInTree)

	minNode := ds.TreeMinimum(root)
	assert.Equal(t, &n0, minNode)
	maxNode := ds.TreeMaximum(root)
	assert.Equal(t, &n3, maxNode)
}

func TestBTree(t *testing.T) {
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

	ans, index := bTree.Search(4)
	assert.Equal(t, bRoot, ans)
	assert.Equal(t, 0, index)

	bTree.Insert(-4)
	insertedChild := bTree.Root.GetChildren()[0]
	expectedKey := []int{-4, -1, 1, 3}
	assert.Equal(t, 4, insertedChild.NumberOfNodes())
	assert.Equal(t, expectedKey, insertedChild.GetKeyList())
}
