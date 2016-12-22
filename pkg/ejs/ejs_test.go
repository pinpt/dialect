package ejs

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/ejs"
)

func TestEJS(t *testing.T) {
	reader := strings.NewReader(`
<div>
	<%= foo %>
</div>
`)
	result, err := dialect.Examine("EJS", "foo.ejs", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 5 {
		t.Fatalf("result.Loc should have been 5, was %d", result.Loc)
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

func TestEJSWithJSMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`
<div>
	<%
		/*
	    this is a embedded comment
		 */
		 var a = 1
	 %>
</div>
`)
	result, err := dialect.Examine("EJS", "foo.ejs", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 10 {
		t.Fatalf("result.Loc should have been 10, was %d", result.Loc)
	}
	if result.Sloc != 5 {
		t.Fatalf("result.Sloc should have been 5, was %d", result.Sloc)
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

func TestEJSWithJSSingleLineComment(t *testing.T) {
	reader := strings.NewReader(`
<div>
	<%
		 // this is a embedded comment
		 var a = 1
		 var b = 2
		 var c = 3
	 %>
</div>
`)
	result, err := dialect.Examine("EJS", "foo.ejs", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 10 {
		t.Fatalf("result.Loc should have been 10, was %d", result.Loc)
	}
	if result.Sloc != 7 {
		t.Fatalf("result.Sloc should have been 7, was %d", result.Sloc)
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
