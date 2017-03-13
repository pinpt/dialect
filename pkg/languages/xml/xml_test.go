package xml

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect/pkg/implementation"
	dialect "github.com/pinpt/dialect/pkg/types"
)

func init() {
	dialect.RegisterDialectImplementation(implementation.New())
}

func TestXMLNoComments(t *testing.T) {
	reader := strings.NewReader(`<?xml version="1.0"?>
<root/>
`)
	result, err := dialect.Examine("XML", "foo.xml", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 3 {
		t.Fatalf("result.Loc should have been 3, was %d", result.Loc)
	}
	if result.Sloc != 2 {
		t.Fatalf("result.Sloc should have been 2, was %d", result.Sloc)
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

func TestXMLSingleLineComment(t *testing.T) {
	reader := strings.NewReader(`<?xml version="1.0"?>
<root>
	<!-- should be empty -->
</root>
`)
	result, err := dialect.Examine("XML", "foo.xml", reader, nil)
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
	if result.Comments != 1 {
		t.Fatalf("result.Comments should have been 1, was %d", result.Comments)
	}
	if result.Blanks != 1 {
		t.Fatalf("result.Blanks should have been 1, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestXMLMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`<?xml version="1.0"?>
<root>
	<!--
		should be empty
	-->
</root>
`)
	result, err := dialect.Examine("XML", "foo.xml", reader, nil)
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
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 3, was %d", result.Comments)
	}
	if result.Blanks != 1 {
		t.Fatalf("result.Blanks should have been 1, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestHTMLMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`<!DOCTYPE html>
<html lang="en">
<head>
</head>
<body>
	<div>
		<!--
			empty line
		-->
	</div>
</body>
</html>
`)
	result, err := dialect.Examine("HTML", "foo.html", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 13 {
		t.Fatalf("result.Loc should have been 13, was %d", result.Loc)
	}
	if result.Sloc != 9 {
		t.Fatalf("result.Sloc should have been 9, was %d", result.Sloc)
	}
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 3, was %d", result.Comments)
	}
	if result.Blanks != 1 {
		t.Fatalf("result.Blanks should have been 1, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestXHTMLMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`<?xml version="1.0"?>
<html lang="en">
<head>
</head>
<body>
	<div>
		<!--
			empty line
		-->
	</div>
</body>
</html>
`)
	result, err := dialect.Examine("XHTML", "foo.xhtml", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 13 {
		t.Fatalf("result.Loc should have been 13, was %d", result.Loc)
	}
	if result.Sloc != 9 {
		t.Fatalf("result.Sloc should have been 9, was %d", result.Sloc)
	}
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 3, was %d", result.Comments)
	}
	if result.Blanks != 1 {
		t.Fatalf("result.Blanks should have been 1, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
