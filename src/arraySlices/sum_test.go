package arraySlices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {

		number := []int{1, 2, 3}

		got := Sum(number)
		want := 6

		assert.Equal(t, got, want)
	})
}

func TestSumAll(t *testing.T) {

	t.Run("all ", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assert.Equal(t, want, got)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assert.Equal(t, want, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assert.Equal(t, want, got)
	})
}
