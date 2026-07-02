package pythonfmt

import "testing"

func TestPythonFormatter(t *testing.T) {
	t.Parallel()

	formatter := DefaultPythonFormatter

	src := []byte("def add(a,b):\n if a>b:\n   return a\n return b\n")
	got, err := formatter.Format(src)
	if err != nil {
		t.Fatalf("could not format the source: %v", err)
	}

	expected := "def add(a, b):\n    if a > b:\n        return a\n    return b\n"
	if string(got) != expected {
		t.Errorf("expected %q, got %q", expected, string(got))
	}
}

func TestPythonFormatterInvalidSyntax(t *testing.T) {
	t.Parallel()

	formatter := DefaultPythonFormatter

	src := []byte("def add(a,b):\n if a>b\n   return a\n")
	_, err := formatter.Format(src)
	if err == nil {
		t.Fatal("expected an error for invalid Python syntax")
	}
}
