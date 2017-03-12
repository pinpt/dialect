package docker

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
)

func TestDockerfile(t *testing.T) {
	reader := strings.NewReader(`# this a comment
ADD env test`)
	result, err := dialect.Examine("Dockerfile", "Dockerfile", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 2 {
		t.Fatalf("result.Loc should have been 2, was %d", result.Loc)
	}
	if result.Sloc != 1 {
		t.Fatalf("result.Sloc should have been 1, was %d", result.Sloc)
	}
	if result.Comments != 1 {
		t.Fatalf("result.Comments should have been 1, was %d", result.Comments)
	}
	if result.Blanks != 0 {
		t.Fatalf("result.Blanks should have been 0, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
