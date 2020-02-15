package algorithms

import (
	"errors"
)

func QuickSort(list []int) ([]int, error) {
	unsorted := make([]int, len(list))
	copy(unsorted, list)
	if unsorted == nil {
		return []int{}, errors.New("nil")
	}
	return quickSort(unsorted, 0, len(unsorted)-1), nil
}

func partition(list []int, p, r int) int {
	x := list[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if list[j] <= x {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[r] = list[r], list[i+1]
	return i + 1
}

func quickSort(list []int, p, r int) []int {
	if p < r {
		q := partition(list, p, r)
		quickSort(list, p, q-1)
		quickSort(list, q+1, r)
	}
	return list
}
