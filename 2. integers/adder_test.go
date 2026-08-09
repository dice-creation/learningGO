package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	t.Run("Add 2 + 2", func (t *testing.T){
		sum := Add(2, 2)
		expected := 4
		assertExpectedSum(t, expected, sum)
	})
}


func assertExpectedSum(t testing.TB, expected, sum int) {
	if expected != sum {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func ExampleAdd(){
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}