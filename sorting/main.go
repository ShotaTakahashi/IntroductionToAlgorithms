package main

import (
	"fmt"

	alg "github.com/ShotaTakahashi/IntroductionToAlgorithms/sorting/algorithms"
	ds "github.com/ShotaTakahashi/IntroductionToAlgorithms/sorting/dataStructure"
)

func main() {
	list := []int{10, 4, 8, 12, 2}
	sorted, _ := alg.InsertionSort(list)
	fmt.Println(sorted)

	list = []int{22, 32, 8, 64, 42, 53, 99}
	sorted, _ = alg.HeapSort(list)
	fmt.Println(sorted)

	list = []int{22, 32, 8, 64, 42, 53, 99}
	sorted, _ = alg.QuickSort(list)
	fmt.Println(sorted)

	n0 := ds.Node{Key: 2, Left: nil, Right: nil}
	n1 := ds.Node{Key: 5, Left: nil, Right: nil}
	n2 := ds.Node{Key: 5, Left: &n0, Right: &n1}
	n3 := ds.Node{Key: 8, Left: nil, Right: nil}
	n4 := ds.Node{Key: 7, Left: nil, Right: &n3}
	root := ds.Node{Key: 6, Left: &n2, Right: &n4}
	fmt.Println(ds.IsBinarySearchTree(root))
	fmt.Println(ds.TreeSearch(root, 3))
}
