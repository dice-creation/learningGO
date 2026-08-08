package arrays

import "testing"
import "slices"

func TestSum(t *testing.T){
	t.Run("collection of 5 numbers", func (t *testing.T){
		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		assertSum(t, got, want, numbers)
	})

	t.Run("collection of any size", func (t *testing.T){
		numbers := []int{1,2,3}

		got := Sum(numbers)
		want := 6
		assertSum(t, got,want, numbers)
	})

	t.Run("Sum all", func (t *testing.T) {
		a := []int{1,2}
		b := []int{0,9}

		got := SumAll(a, b)
		want := []int{3,9}

		if !slices.Equal(got,want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}


func assertSum(t testing.TB, got, want int, numbers []int){
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}