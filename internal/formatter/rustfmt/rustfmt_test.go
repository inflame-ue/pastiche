package rustfmt

import "testing"

func TestRustFormat(t *testing.T) {
	t.Parallel()

	formatter := DefaultRustFormatter

	src := []byte("fn main() {\nprintln!(\"hello\");\n}")
	got, err := formatter.Format(src)
	if err != nil {
		t.Fatalf("could not format the source: %v", err)
	}

	expected := "fn main() {\n    println!(\"hello\");\n}\n"
	if string(got) != expected {
		t.Errorf("expected %q, got %q", expected, string(got))
	}
}

func TestInvalidRustFormat(t *testing.T) {
	t.Parallel()

	formatter := DefaultRustFormatter

	src := []byte("fn main() {\n x = \n}")
	_, err := formatter.Format(src)
	if err == nil {
		t.Fatal("expected an error for invalid Rust syntax")
	}
}
