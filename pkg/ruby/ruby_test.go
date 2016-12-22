package ruby

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/ruby"
)

func TestRubyNoComments(t *testing.T) {
	reader := strings.NewReader("puts \"hello\"")
	result, err := dialect.Examine("Ruby", "foo.rb", reader, nil)
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

func TestRubySingleLineComments(t *testing.T) {
	reader := strings.NewReader(`
# comment
puts \"hello\"`)
	result, err := dialect.Examine("Ruby", "foo.rb", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 3 {
		t.Fatalf("result.Loc should have been 3, was %d", result.Loc)
	}
	if result.Sloc != 1 {
		t.Fatalf("result.Sloc should have been 1, was %d", result.Sloc)
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

func TestRubyMultiLineComments(t *testing.T) {
	reader := strings.NewReader(`
=begin
this is a multiple
line comment
=end
puts \"hello\"`)
	result, err := dialect.Examine("Ruby", "foo.rb", reader, nil)
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
	if result.Comments != 4 {
		t.Fatalf("result.Comments should have been 4, was %d", result.Comments)
	}
	if result.Blanks != 1 {
		t.Fatalf("result.Blanks should have been 1, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestRubySelenium(t *testing.T) {
	// from https://wiki.saucelabs.com/display/DOCS/Ruby+Test+Setup+Example
	reader := strings.NewReader(`require "selenium/webdriver"

caps = {
  :platform => "Windows 7",
  :browserName => "Chrome",
  :version => "45"
}

driver = Selenium::WebDriver.for(:remote,
    :url => "https://YOUR_SAUCE_USERNAME:YOUR_SAUCE_ACCESS_KEY@ondemand.saucelabs.com:443/wd/hub",
    :desired_capabilities => caps)

driver.get('http://saucelabs.com/test/guinea-pig')

puts "title of webpage is: #{driver.title()}"

driver.quit()`)
	result, err := dialect.Examine("Ruby", "foo.rb", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 17 {
		t.Fatalf("result.Loc should have been 17, was %d", result.Loc)
	}
	if result.Sloc != 12 {
		t.Fatalf("result.Sloc should have been 12, was %d", result.Sloc)
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
