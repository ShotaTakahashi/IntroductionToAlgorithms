package algorithms

import "errors"

func InsertionSort(list []int) ([]int, error) {
	if list == nil {
		return []int{}, errors.New("nil")
	}
	for j := 1; j < len(list); j++ {
		key := list[j]
		i := j-1
		for i >= 0 && list[i] > key {
			list[i+1] = list[i]
			i -= 1
		}
		list[i+1] = key
	}
	return list, nil
}