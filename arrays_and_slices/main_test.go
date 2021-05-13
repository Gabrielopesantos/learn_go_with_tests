package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum all test", func(t *testing.T) {
		got := SumAll([]int{2, 3}, []int{9, 0})
		want := []int{5, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sum all tails test", func(t *testing.T) {
		got := SumAllTails([]int{1, 3}, []int{9, 0})
		want := []int{3, 0}

		checkSums(t, got, want)
	})

	t.Run("Sum all tails test empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(t, got, want)
	})
}
