package dialect

import (
	"strings"
	"testing"
)

func TestLoadDialect(t *testing.T) {
	config := CreateDefaultConfiguration()
	result, err := Examine("JavaScript", "test.js", strings.NewReader("var a = 1"), config)
	if err != nil {
		t.Fatal(err)
	}
	if result.Blanks != 0 {
		t.Fatalf("expected Blanks to be 0 was %d", result.Blanks)
	}
	if result.Comments != 0 {
		t.Fatalf("expected Comments to be 0 was %d", result.Comments)
	}
	if result.Sloc != 1 {
		t.Fatalf("expected Sloc to be 1 was %d", result.Sloc)
	}
	if result.Loc != 1 {
		t.Fatalf("expected Loc to be 1 was %d", result.Loc)
	}
	if result.IsTest {
		t.Fatalf("expected Test to be false was %v", result.IsTest)
	}
}
