package arrnsl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %d want %d, given %d", got, want, numbers)
		}
	})

	t.Run("Collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d want %d, given %d", got, want, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3})
	fmt.Println(sum)
	// output: 6
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %d want %d", got, want)
	}

}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3, 4, 5}, []int{4, 5, 6})
	}
}

func ExampleSumAll() {
	sum := SumAll([]int{1, 2, 3}, []int{3, 4, 5}, []int{5, 6, 7})
	fmt.Println(sum)
	// output: [6 12 18]
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %d want %d", got, want)
		}
	}
	t.Run("a sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(t, got, want)
	})
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3, 4, 5}, []int{4, 5, 6})
	}
}

func ExampleSumAllTails() {
	sum := SumAllTails([]int{1, 2, 3}, []int{3, 4, 5}, []int{5, 6, 7})
	fmt.Println(sum)
	// output: [5 9 13]
}
