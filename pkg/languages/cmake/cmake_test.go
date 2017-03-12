package cmake

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
)

func TestCMakeNoComments(t *testing.T) {
	reader := strings.NewReader("install(TARGETS main.x DESTINATION appcelerator/ingot)")
	result, err := dialect.Examine("CMake", "foo.cmake", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 1 {
		t.Fatalf("result.Loc should have been 1, was %d", result.Loc)
	}
	if result.Sloc != 1 {
		t.Fatalf("result.Sloc should have been 1, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 0 {
		t.Fatalf("result.Blanks should have been 0, was %d", result.Blanks)
	}
}

func TestCMakeSingleLineComment(t *testing.T) {
	reader := strings.NewReader(`
# this is a comment
install(TARGETS main.x DESTINATION appcelerator/ingot)`)
	result, err := dialect.Examine("CMake", "foo.cmake", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 3 {
		t.Fatalf("result.Loc should have been 3, was %d", result.Loc)
	}
	if result.Sloc != 1 {
		t.Fatalf("result.Sloc should have been 1, was %d", result.Sloc)
	}
	if result.Comments != 1 {
		t.Fatalf("result.Comments should have been 1, was %d", result.Comments)
	}
	if result.Blanks != 1 {
		t.Fatalf("result.Blanks should have been 1, was %d", result.Blanks)
	}
}
