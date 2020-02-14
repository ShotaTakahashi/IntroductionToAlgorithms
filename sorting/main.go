package main

import "fmt"

func main()  {
	list := []int{10, 4, 8, 12, 2}
	sorted := InsertionSort(list)
	fmt.Printf(sorted)
}
