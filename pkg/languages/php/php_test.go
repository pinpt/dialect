package php

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect/pkg/implementation"
	dialect "github.com/pinpt/dialect/pkg/types"
)

func init() {
	dialect.RegisterDialectImplementation(implementation.New())
}

func TestPHPSingleLine(t *testing.T) {
	reader := strings.NewReader("<?php ?>")
	result, err := dialect.Examine("PHP", "foo.php", reader, nil)
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

func TestPHPWebdriverTest(t *testing.T) {
	// from https://wiki.saucelabs.com/display/DOCS/PHP+Test+Setup+Example
	reader := strings.NewReader(`<?php
	// Setup: $ php composer.phar require facebook/webdriver

	require_once('vendor/autoload.php');
	use Facebook\WebDriver\Remote\RemoteWebDriver;
	use Facebook\WebDriver\WebDriverBy;

	$web_driver = RemoteWebDriver::create(
	 "https://YOUR_SAUCE_USERNAME:YOUR_SAUCE_ACCESS_KEY@ondemand.saucelabs.com:443/wd/hub",
	 array("platform"=>"Windows 7", "browserName"=>"chrome", "version"=>"40")
	);
	$web_driver->get("https://saucelabs.com/test/guinea-pig");

	/*
	 Test actions here...
	*/

	$web_driver->quit();
?>`)
	result, err := dialect.Examine("PHP", "foo.php", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 19 {
		t.Fatalf("result.Loc should have been 19, was %d", result.Loc)
	}
	if result.Sloc != 11 {
		t.Fatalf("result.Sloc should have been 11, was %d", result.Sloc)
	}
	if result.Comments != 4 {
		t.Fatalf("result.Comments should have been 4, was %d", result.Comments)
	}
	if result.Blanks != 4 {
		t.Fatalf("result.Blanks should have been 4, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}
