package json

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect/pkg/implementation"
	dialect "github.com/pinpt/dialect/pkg/types"
)

func init() {
	dialect.RegisterDialectImplementation(implementation.New())
}

func TestJSONWithBlankLines(t *testing.T) {
	reader := strings.NewReader(`{
	"a": "b",

	"b": 1
}`)
	result, err := dialect.Examine("JSON", "foo.json", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 5 {
		t.Fatalf("result.Loc should have been 1, was %d", result.Loc)
	}
	if result.Sloc != 4 {
		t.Fatalf("result.Sloc should have been 4, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 1 {
		t.Fatalf("result.Blanks should have been 1, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
