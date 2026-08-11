package dictionary

import "testing"

func TestSearch(t *testing.T) {
	// t.Run("Searching", func(t *testing.T) {
	// dictionary := map[string]string{"test": "this is just a test"}
	// got := Search(dictionary, "test")
	// want := "this is just a test"

	// assertStrings(t, got, want)
	// })

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("did not expect an error")
		}
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, got := dictionary.Search("Unkown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, ErrNotFound)
	})
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
