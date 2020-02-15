package algorithms

import (
	"errors"
)

func InsertionSort(list []int) ([]int, error) {
	unsorted := make([]int, len(list))
	copy(unsorted, list)
	if unsorted == nil {
		return []int{}, errors.New("nil")
	}
	for j := 1; j < len(unsorted); j++ {
		key := unsorted[j]
		i := j - 1
		for i >= 0 && unsorted[i] > key {
			unsorted[i+1] = unsorted[i]
			i -= 1
		}
		unsorted[i+1] = key
	}
	return unsorted, nil
}
