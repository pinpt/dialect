package r

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect/pkg/implementation"
	dialect "github.com/pinpt/dialect/pkg/types"
)

func init() {
	dialect.RegisterDialectImplementation(implementation.New())
}

func TestR(t *testing.T) {
	reader := strings.NewReader(`"This function takes a value x, and does things and returns things that take several lines to explain"
doEverythingOften <- function(x) {
  # Non! Comment it out! We'll just do it once for now.
  "if (x %in% 1:9) {
		 doTenEverythings()
  }"
  doEverythingOnce()
  return(list(
		everythingDone = TRUE,
		howOftenDone = 1
  ))
}`)
	result, err := dialect.Examine("R", "test.r", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 12 {
		t.Fatalf("result.Loc should have been 12, was %d", result.Loc)
	}
	if result.Sloc != 7 {
		t.Fatalf("result.Sloc should have been 7, was %d", result.Sloc)
	}
	if result.Comments != 5 {
		t.Fatalf("result.Comments should have been 5, was %d", result.Comments)
	}
	if result.Blanks != 0 {
		t.Fatalf("result.Blanks should have been 0, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
