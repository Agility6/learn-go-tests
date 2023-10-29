package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestExampleClone(t *testing.T) {
	clone := ExampleClone("abc")
	expected := "abc"

	if clone != expected {
		t.Errorf("expected %q but got %q", expected, clone)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// it runs b.N times and measures how long it takes.
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
