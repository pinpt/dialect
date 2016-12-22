package makefile

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/makefile"
)

func TestMakefile(t *testing.T) {
	reader := strings.NewReader(`.phony foo

# this is a comment

foo:
	@echo hello
`)
	result, err := dialect.Examine("Makefile", "Makefile", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 7 {
		t.Fatalf("result.Loc should have been 7, was %d", result.Loc)
	}
	if result.Sloc != 3 {
		t.Fatalf("result.Sloc should have been 3, was %d", result.Sloc)
	}
	if result.Comments != 1 {
		t.Fatalf("result.Comments should have been 1, was %d", result.Comments)
	}
	if result.Blanks != 3 {
		t.Fatalf("result.Blanks should have been 3, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
