package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Repeat character 2 times", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	text := Repeat("a", 5)
	fmt.Println(text)
	// Output: aaaaa
}
