package main

import (
	"fmt"
	alg "github.com/ShotaTakahashi/IntroductionToAlgorithms/sorting/algorithms"
)

func main()  {
	list := []int{10, 4, 8, 12, 2}
	sorted := alg.InsertionSort(list)
	fmt.Println(sorted)
}
