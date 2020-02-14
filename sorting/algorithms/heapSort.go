package algorithms

import "errors"

func HeapSort(list []int) ([]int, error) {
	if list == nil {
		return []int{}, errors.New("nil")
	}
	maxHeap := buildMaxHeap(list)
	size := len(maxHeap) - 1
	for i := size; i > 0; i-- {
		maxHeap[0], maxHeap[i] = maxHeap[i], maxHeap[0]
		size--
		maxHeap = maxHeapify(maxHeap, 0, size)
	}
	return maxHeap, nil
}

func buildMaxHeap(list []int) (result []int) {
	for i := len(list) / 2; i > -1; i-- {
		result = maxHeapify(list, i, len(list) - 1)
	}
	return result
}

func maxHeapify(list []int, i, size int) []int {
	left := left(i)
	right := right(i)
	var largest int
	if left <= size && list[left] > list[i] {
		largest = left
	} else {
		largest = i
	}

	if right <= size && list[right] > list[largest] {
		largest = right
	}

	if largest != i {
		list[i], list[largest] = list[largest], list[i]
		return maxHeapify(list, largest, size)
	}
	return list
}

func left(i int) int {
	return 2 * i + 1
}

func right(i int) int {
	return 2 * i
}