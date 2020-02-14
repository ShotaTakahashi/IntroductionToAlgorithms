package test

import (
	alg "github.com/ShotaTakahashi/IntroductionToAlgorithms/sorting/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T)  {
	unsorted := []int{10, 4, 8, 12, 2, 11}
	sorted, _ := alg.InsertionSort(unsorted)
	expected := []int{2, 4, 8, 10, 11, 12}
	assert.Equal(t, expected, sorted)

	unsorted = nil
	sorted, _ = alg.InsertionSort(nil)
	assert.Equal(t, []int{}, sorted)
}
