package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T){

	t.Run("Repeat A, 5 times", func (t *testing.T){
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertRepeat(t, repeated, expected)
	})
	
	t.Run("Testing with added Count", func (t *testing.T){
		repeated := Repeat("c", 3)
		expected := "ccc"

		assertRepeat(t, repeated, expected)
	})

	t.Run("Testing with 0 repeatCount", func (t *testing.T){
		repeated := Repeat("a", 0)
		expected := ""

		assertRepeat(t, repeated, expected)
	})
}



func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 4)
	}
}


func assertRepeat(t testing.TB, repeat, expected string){
	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 7)
	fmt.Println(repeated)
	// Output: aaaaaaa
}
