package main

import (
	"fmt"

	alg "github.com/ShotaTakahashi/IntroductionToAlgorithms/sorting/algorithms"
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
}
