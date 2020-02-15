package test

import (
	"testing"

	ds "github.com/ShotaTakahashi/IntroductionToAlgorithms/sorting/dataStructure"
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

	_, isInTree := ds.TreeSearch(root, 3)
	assert.Equal(t, false, isInTree)

	node, isInTree := ds.TreeSearch(root, 5)
	assert.Equal(t, &n2, node)
	assert.Equal(t, true, isInTree)
}
