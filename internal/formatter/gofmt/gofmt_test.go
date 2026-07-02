package gofmt

import "testing"

func TestGoFormat(t *testing.T) {
	t.Parallel()

	formatter := NewGoFormatter()

	src := []byte("func main(){\na:=1\nb:=2\n_=a+b}")
	got, err := formatter.Format(src)
	if err != nil {
		t.Fatalf("could not format the source: %v", err)
	}

	expected := "func main() {\n	a := 1\n	b := 2\n	_ = a + b\n}"
	if string(got) != expected {
		t.Errorf("expected %q, got %q", expected, string(got))
	}
}

func TestInvalidGoFormat(t *testing.T) {
	t.Parallel()

	formatter := NewGoFormatter()

	src := []byte("func main()\na:=1\nb:=2\n_=a+b}")
	_, err := formatter.Format(src)
	if err == nil {
		t.Error("expected an error for invalid Go syntax")
	}
}
