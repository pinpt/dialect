package def

import (
	"errors"
	"strings"
	"testing"

	"github.com/pinpt/dialect"
)

func TestDefault(t *testing.T) {
	reader := strings.NewReader("a is 1")
	result, err := dialect.Examine("SomeUnknownLanguage", "foo.bar", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 1 {
		t.Fatalf("result.Loc should have been 1, was %d", result.Loc)
	}
	if result.Sloc != 0 {
		t.Fatalf("result.Sloc should have been 0, was %d", result.Sloc)
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

func TestDefaultCallback(t *testing.T) {
	var invoked bool
	callback := func(language string, line *dialect.DialectLine) error {
		invoked = true
		if line.IsComment {
			return errors.New("IsComment should have been false")
		}
		if line.IsBlank {
			return errors.New("IsBlank should have been false")
		}
		if line.IsCode {
			return errors.New("IsCode should have been false")
		}
		if line.LineNumber != 1 {
			return errors.New("LineNumber should have been 1")
		}
		return nil
	}
	reader := strings.NewReader("a is 1")
	result, err := dialect.Examine("SomeUnknownLanguage", "foo.bar", reader, dialect.CreateConfigurationWithCallback(callback))
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if invoked == false {
		t.Fatal("callback was not invoked")
	}
	if result.Loc != 1 {
		t.Fatalf("result.Loc should have been 1, was %d", result.Loc)
	}
	if result.Sloc != 0 {
		t.Fatalf("result.Sloc should have been 0, was %d", result.Sloc)
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

func TestDefaultCallbackWithError(t *testing.T) {
	callback := func(language string, line *dialect.DialectLine) error {
		return errors.New("test")
	}
	reader := strings.NewReader("a is 1")
	result, err := dialect.Examine("SomeUnknownLanguage", "foo.bar", reader, dialect.CreateConfigurationWithCallback(callback))
	if err == nil {
		t.Fatal("error should have been returned")
	}
	if result != nil {
		t.Fatal("result should have been nil")
	}
	if err.Error() != "test" {
		t.Fatal("error should have been \"test\"")
	}
}
