package cstyle

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect/pkg/implementation"
	dialect "github.com/pinpt/dialect/pkg/types"
)

func init() {
	dialect.RegisterDialectImplementation(implementation.New())
}

func TestLanguageCWithOneLine(t *testing.T) {
	reader := strings.NewReader("int a")
	result, err := dialect.Examine("C", "foo.c", reader, nil)
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
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestLanguageCWithMultipleLinesSingleLineComment(t *testing.T) {
	reader := strings.NewReader(`
// example
int main() {
	int a = 1
	return a
}
`)
	result, err := dialect.Examine("C", "foo.c", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 7 {
		t.Fatalf("result.Loc should have been 7, was %d", result.Loc)
	}
	if result.Sloc != 4 {
		t.Fatalf("result.Sloc should have been 4, was %d", result.Sloc)
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

func TestLanguageCWithMultipleLinesMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`
/*
 hello
 */
int main() {
	int a = 1
	return a
}
`)
	result, err := dialect.Examine("C", "foo.c", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 9 {
		t.Fatalf("result.Loc should have been 9, was %d", result.Loc)
	}
	if result.Sloc != 4 {
		t.Fatalf("result.Sloc should have been 4, was %d", result.Sloc)
	}
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 3, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestLanguageCWithSingleLineMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`
/** hello */
int main() {
	int a = 1
	return a
}
`)
	result, err := dialect.Examine("C", "foo.c", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 7 {
		t.Fatalf("result.Loc should have been 7, was %d", result.Loc)
	}
	if result.Sloc != 4 {
		t.Fatalf("result.Sloc should have been 4, was %d", result.Sloc)
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

func TestLanguageCPlusPlusWithSingleLineMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`
/** hello */
int main() {
	int a = 1
	return a
}
`)
	result, err := dialect.Examine("C++", "foo.cpp", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 7 {
		t.Fatalf("result.Loc should have been 7, was %d", result.Loc)
	}
	if result.Sloc != 4 {
		t.Fatalf("result.Sloc should have been 4, was %d", result.Sloc)
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

func TestLanguageCPlusPlusWithSingleMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`
/*
hello */
int main() {
	int a = 1
	return a
}
`)
	result, err := dialect.Examine("C++", "foo.cpp", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 8 {
		t.Fatalf("result.Loc should have been 8, was %d", result.Loc)
	}
	if result.Sloc != 4 {
		t.Fatalf("result.Sloc should have been 4, was %d", result.Sloc)
	}
	if result.Comments != 2 {
		t.Fatalf("result.Comments should have been 1, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
