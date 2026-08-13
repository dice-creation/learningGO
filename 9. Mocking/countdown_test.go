package mocking

import "testing"
import "bytes"

func TestCountdown(t *testing.T){
	t.Run("counting", func (t *testing.T){
		buffer := &bytes.Buffer{}

		Countdown(buffer)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}