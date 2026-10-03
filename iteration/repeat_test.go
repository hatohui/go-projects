package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 9)
	expected := "aaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 1)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 4)
	fmt.Println(repeated)
	// Output: "aaaa"
}
