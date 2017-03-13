package swift

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect/pkg/implementation"
	dialect "github.com/pinpt/dialect/pkg/types"
)

func init() {
	dialect.RegisterDialectImplementation(implementation.New())
}

func TestSwift(t *testing.T) {
	reader := strings.NewReader(`class NamedShape {
    var numberOfSides: Int = 0
    var name: String

    init(name: String) {
        self.name = name
    }
	 // example
    func simpleDescription() -> String {
        return "A shape with \(numberOfSides) sides."
    }
}
`)
	result, err := dialect.Examine("Swift", "foo.swift", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 13 {
		t.Fatalf("result.Loc should have been 13, was %d", result.Loc)
	}
	if result.Sloc != 10 {
		t.Fatalf("result.Sloc should have been 10, was %d", result.Sloc)
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

func TestSwiftTestCase(t *testing.T) {
	reader := strings.NewReader(`import XCTest
import SwiftFonts

class FontSorterTests: XCTestCase {

    let sorter = FontSorter()

    func testCompareHyphenWithNoHyphen() {
        let fonts = ["Arial-ItalicMT", "ArialMT"]
        let expected = ["ArialMT", "Arial-ItalicMT"]
        let sorted = sorter.sortFontNames(fonts)
        XCTAssertEqual(expected[0], sorted[0], "the array should be sorted properly")
        XCTAssertEqual(expected[1], sorted[1], "the array should be sorted properly")
    }

    func testCompareHyphenWithHyphen() {
        let fonts = ["Avenir-Roman", "Avenir-Oblique"]
        let expected = ["Avenir-Oblique", "Avenir-Roman"]
        let sorted = sorter.sortFontNames(fonts)
        XCTAssertEqual(expected[0], sorted[0], "when two fonts contain a hyphen, they should be sorted alphabetically")
        XCTAssertEqual(expected[1], sorted[1], "when two fonts contain a hyphen, they should be sorted alphabetically")
    }
}
`)
	result, err := dialect.Examine("Swift", "foo.swift", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 24 {
		t.Fatalf("result.Loc should have been 24, was %d", result.Loc)
	}
	if result.Sloc != 19 {
		t.Fatalf("result.Sloc should have been 19, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 5 {
		t.Fatalf("result.Blanks should have been 5, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}
