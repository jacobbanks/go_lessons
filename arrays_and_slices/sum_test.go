package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3}
	got := AddNums(nums)
	want := 6
	if got != want {
		t.Fatalf("got %d wanted %d, given %v", got, want, nums)
	}
}

func TestSumAll(t *testing.T) {
	numOne := []int{1, 2, 3}
	numTwo := []int{3, 2, 1}

	got := SumAll(numOne, numTwo)
	want := []int{6, 6}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v wanted %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	assertSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v wanted %v", got, want)
		}	
	}
	t.Run("Add Tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		assertSums(t, got, want)
	})

	t.Run("Add Tails, allowing for empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5, 6})
		want := []int{0, 11}
		assertSums(t, got, want)
	})
}
