package python

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/python"
)

func TestPythonNoComments(t *testing.T) {
	reader := strings.NewReader("puts \"hello\"")
	result, err := dialect.Examine("Python", "foo.py", reader, nil)
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

func TestPythonSingleLineComment(t *testing.T) {
	reader := strings.NewReader(`
# this is a comment
puts \"hello\"`)
	result, err := dialect.Examine("Python", "foo.py", reader, nil)
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

func TestPythonMultiLineComment(t *testing.T) {
	reader := strings.NewReader(`
"""
this is a comment
that spans multiple lines
"""
puts \"hello\"`)
	result, err := dialect.Examine("Python", "foo.py", reader, nil)
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

func TestPythonSelenium(t *testing.T) {
	// from https://wiki.saucelabs.com/display/DOCS/Python+Test+Setup+Example
	reader := strings.NewReader(`from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities

# This is the only code you need to edit in your existing scripts.
# The command_executor tells the test to run on Sauce, while the desired_capabilities
# parameter tells us which browsers and OS to spin up.
desired_cap = {
    'platform': "Mac OS X 10.9",
    'browserName': "chrome",
    'version': "31",
}
driver = webdriver.Remote(
   command_executor='http://YOUR_SAUCE_USERNAME:YOUR_SAUCE_ACCESS_KEY@ondemand.saucelabs.com:80/wd/hub',
   desired_capabilities=desired_cap)

# This is your test logic. You can add multiple tests here.
driver.implicitly_wait(10)
driver.get("http://www.google.com")
if not "Google" in driver.title:
    raise Exception("Unable to load google page!")
elem = driver.find_element_by_name("q")
elem.send_keys("Sauce Labs")
elem.submit()
print driver.title

# This is where you tell Sauce Labs to stop running tests on your behalf.
# It's important so that you aren't billed after your test finishes.
driver.quit()`)
	result, err := dialect.Examine("Python", "foo.py", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 29 {
		t.Fatalf("result.Loc should have been 29, was %d", result.Loc)
	}
	if result.Sloc != 20 {
		t.Fatalf("result.Sloc should have been 20, was %d", result.Sloc)
	}
	if result.Comments != 6 {
		t.Fatalf("result.Comments should have been 6, was %d", result.Comments)
	}
	if result.Blanks != 3 {
		t.Fatalf("result.Blanks should have been 3, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}
