package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum with fixed size array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of slices", func(t *testing.T) {
		result := SumAll([]int{1, 2, 3}, []int{3, 5})
		expected := []int{6, 8}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %v expected %v", result, expected)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, result, expected []int) {
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %v expected %v", result, expected)
		}
	}

	t.Run("sum tails from slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{2, 9, 2})
		expected := []int{2, 11}

		checkSums(t, result, expected)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{2, 9, 2})
		expected := []int{0, 11}

		checkSums(t, result, expected)
	})
}
