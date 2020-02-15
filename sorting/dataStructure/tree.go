package dataStructure

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n Node) isBinarySearchTree() bool {
	if n.Left == nil {
		return n.Right.Key >= n.Key
	}
	if n.Right == nil {
		return n.Left.Key <= n.Key
	}
	return n.Left.Key <= n.Key && n.Right.Key >= n.Key
}

func IsBinarySearchTree(root Node) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return root.isBinarySearchTree()
	}
	if root.isBinarySearchTree() {
		return IsBinarySearchTree(*root.Left) && IsBinarySearchTree(*root.Right)
	}
	return false
}

func TreeSearch(root Node, key int) (*Node, bool) {
	for &root != nil && key != root.Key {
		if root.Left == nil && root.Right == nil {
			return nil, false
		}
		if key < root.Key {
			root = *root.Left
		} else {
			root = *root.Right
		}
	}
	return &root, true
}
