package gofmt

import (
	"go/format"
	"testing"
)

func TestGoFormatter(t *testing.T) {
	t.Parallel()

	formatter := NewGoFormatter()

	src := `package main
func main() {
x:=1
y:=2
_=x+y
}`
	expected, err := format.Source([]byte(src))
	if err != nil {
		t.Fatalf("could not format the expect string: %v", err)
	}
	got, err := formatter.Format([]byte(src))
	if err != nil {
		t.Fatalf("could not format the got string: %v", err)
	}

	if string(expected) != string(got) {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
