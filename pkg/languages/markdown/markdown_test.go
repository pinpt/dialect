package markdown

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
)

func TestMarkdown(t *testing.T) {
	reader := strings.NewReader(`# Hello
This is a good example of a markdown.

Bye
`)
	result, err := dialect.Examine("Markdown", "foo.md", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 5 {
		t.Fatalf("result.Loc should have been 5 was %d", result.Loc)
	}
	if result.Sloc != 3 {
		t.Fatalf("result.Sloc should have been 3, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
