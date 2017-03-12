package coffeescript

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
)

func TestCoffeeScriptSingleLineComment(t *testing.T) {
	reader := strings.NewReader(`
# Define the main object
ipaddr = {}
`)
	result, err := dialect.Examine("CoffeeScript", "foo.coffee", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 4 {
		t.Fatalf("result.Loc should have been 4, was %d", result.Loc)
	}
	if result.Sloc != 1 {
		t.Fatalf("result.Sloc should have been 1, was %d", result.Sloc)
	}
	if result.Comments != 1 {
		t.Fatalf("result.Comments should have been 1, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestCoffeeScriptMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`
###
this is a multi line comment
###
ipaddr = {}
`)
	result, err := dialect.Examine("CoffeeScript", "foo.coffee", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 6 {
		t.Fatalf("result.Loc should have been 6, was %d", result.Loc)
	}
	if result.Sloc != 1 {
		t.Fatalf("result.Sloc should have been 1, was %d", result.Sloc)
	}
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 1, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
