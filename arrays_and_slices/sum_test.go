package arrays_and_slices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Add collection of any size", func(t *testing.T) {
		nums := []int{1,2,3}
		got := AddNums(nums)
		want := 6
		if got != want {
			t.Fatalf("got %d wanted %d, given %v", got, want, nums)
		}
	})
}
