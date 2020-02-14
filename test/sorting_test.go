package test

import (
	alg "github.com/ShotaTakahashi/IntroductionToAlgorithms/sorting/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSorting(t *testing.T)  {
	cases := []struct {
		unsorted []int
		expected []int
	}{
		{
			unsorted: []int{10, 4, 8, 12, 2, 11},
			expected: []int{2, 4, 8, 10, 11, 12},
		},
		{
				unsorted: []int{22, 32, 8, 64, 42, 53, 99},
				expected: []int{8, 22, 32, 42, 53, 64, 99},
		},
		{
					unsorted: nil,
					expected: []int{},
		},
	}

	for _, c := range cases {
		sorted, _ := alg.InsertionSort(c.unsorted)
		assert.Equal(t, c.expected, sorted)
	}

	for _, c := range cases {
		sorted, _ := alg.HeapSort(c.unsorted)
		assert.Equal(t, c.expected, sorted)
	}
}
