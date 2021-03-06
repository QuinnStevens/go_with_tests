package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
	})
	t.Run("repeat longer string", func(t *testing.T) {
		got := Repeat("blah", 3)
		want := "blahblahblah"

		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
	})
	t.Run("repeat 0 times", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
	})
	t.Run("repeat 7 times", func(t *testing.T) {
		got := Repeat("a", 7)
		want := "aaaaaaa"

		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("blah", 3)
	fmt.Println(res)
	// Output: blahblahblah
}
