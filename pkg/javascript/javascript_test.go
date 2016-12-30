package javascript

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/javascript"
)

func TestJavaScriptSingleLine(t *testing.T) {
	reader := strings.NewReader("var a = 1")
	result, err := dialect.Examine("JavaScript", "foo.js", reader, nil)
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

func TestECMAScriptSingleLine(t *testing.T) {
	reader := strings.NewReader("var a = 1")
	result, err := dialect.Examine("ECMAScript", "foo.js", reader, nil)
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

func TestMochaTest(t *testing.T) {
	reader := strings.NewReader(`
describe("foo", function() {
	it("should be a test", function(done) {
		done()
	});
});
`)
	result, err := dialect.Examine("JavaScript", "foo.js", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 7 {
		t.Fatalf("result.Loc should have been 7, was %d", result.Loc)
	}
	if result.Sloc != 5 {
		t.Fatalf("result.Sloc should have been 5, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}

func TestWebdriverTest(t *testing.T) {
	// from https://wiki.saucelabs.com/display/DOCS/Node.js+Test+Setup+Example
	reader := strings.NewReader(`var webdriver = require('selenium-webdriver'),
    username = "YOUR_SAUCE_USERNAME",
    accessKey = "YOUR_SAUCE_ACCESS_KEY",
    driver;

driver = new webdriver.Builder().
  withCapabilities({
    'browserName': 'chrome',
    'platform': 'Windows XP',
    'version': '43.0',
    'username': username,
    'accessKey': accessKey
  }).
  usingServer("http://" + username + ":" + accessKey +
              "@ondemand.saucelabs.com:80/wd/hub").
  build();

driver.get('http://saucelabs.com/test/guinea-pig');

driver.getTitle().then(function (title) {
    console.log("title is: " + title);
});

driver.quit();
`)
	result, err := dialect.Examine("JavaScript", "foo.js", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 25 {
		t.Fatalf("result.Loc should have been 25, was %d", result.Loc)
	}
	if result.Sloc != 20 {
		t.Fatalf("result.Sloc should have been 20, was %d", result.Sloc)
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

func TestCucumberTest(t *testing.T) {
	// from https://github.com/cucumber/cucumber-js/blob/master/example/example.js
	reader := strings.NewReader(`var featureEditor, stepDefinitionsEditor, $output;

function runFeature() {
  $output.empty();
  $('a[href="#output-tab"]').tab('show');

  var featureSource = featureEditor.getValue();
  var feature = Cucumber.FeatureParser.parse({
    scenarioFilter: new Cucumber.ScenarioFilter({}),
    source: featureSource,
    uri: '/feature'
  });

  Cucumber.clearSupportCodeFns();
  new Function(stepDefinitionsEditor.getValue())();
  var supportCodeLibrary = Cucumber.SupportCodeLibraryBuilder.build({
    cwd: '/',
    fns: Cucumber.getSupportCodeFns()
  });

  var formatterOptions = {
    colorsEnabled: true,
    cwd: '/',
    log: function(data) {
      appendToOutput(ansi_up.ansi_to_html(data));
    },
    supportCodeLibrary: supportCodeLibrary
  };
  var prettyFormatter = Cucumber.FormatterBuilder.build('pretty', formatterOptions);

  var runtime = new Cucumber.Runtime({
    features: [feature],
    listeners: [prettyFormatter],
    supportCodeLibrary: supportCodeLibrary
  });
  return runtime.start();
};

function appendToOutput(data) {
  $output.append(data);
  $output.scrollTop($output.prop("scrollHeight"));
}

function displayError(error) {
  var errorContainer = $('<div>')
  errorContainer.addClass('error').text(error.stack || error);
  appendToOutput(errorContainer)
}

$(function() {
  featureEditor = ace.edit("feature");
  featureEditor.getSession().setMode("ace/mode/gherkin");

  stepDefinitionsEditor = ace.edit("step-definitions");
  stepDefinitionsEditor.getSession().setMode("ace/mode/javascript");

  $output = $('#output');

  window.onerror = displayError;

  $('#run-feature').click(function() {
    runFeature().then(function(success) {
      var exitStatus = success ? '0' : '1';
      var exitStatusContainer = $('<div>');
      exitStatusContainer.addClass('exit-status').text('Exit Status: ' + exitStatus);
      appendToOutput(exitStatusContainer);
    }).catch(displayError);
  });
});
`)
	result, err := dialect.Examine("JavaScript", "foo.js", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 70 {
		t.Fatalf("result.Loc should have been 70, was %d", result.Loc)
	}
	if result.Sloc != 57 {
		t.Fatalf("result.Sloc should have been 57, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 13 {
		t.Fatalf("result.Blanks should have been 13, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}

func TestAvaTest(t *testing.T) {
	// from https://github.com/avajs/ava
	reader := strings.NewReader(`import test from 'ava';

test('foo', t => {
    t.pass();
});

test('bar', async t => {
    const bar = Promise.resolve('bar');

    t.is(await bar, 'bar');
});
`)
	result, err := dialect.Examine("JavaScript", "foo.js", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 12 {
		t.Fatalf("result.Loc should have been 12, was %d", result.Loc)
	}
	if result.Sloc != 8 {
		t.Fatalf("result.Sloc should have been 8, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 4 {
		t.Fatalf("result.Blanks should have been 4, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}

func TestCopyright(t *testing.T) {
	reader := strings.NewReader(`// Copyright (c) 2012 A Pinpoint PBC
var a = 1
`)
	config := dialect.CreateDefaultConfiguration()
	config.DetectCopyrights = true
	result, err := dialect.Examine("JavaScript", "foo.js", reader, config)
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
	if result.Copyrights == nil {
		t.Fatal("result.Copyrights was nil")
	}
	if len(result.Copyrights) != 1 {
		t.Fatalf("result.Copyrights count should have been 1, was %d", len(result.Copyrights))
	}
	copyright := result.Copyrights[0]
	if len(copyright.Dates) != 1 {
		t.Fatalf("result.Copyrights[0].Dates[0] len should have been 1, was %d", len(copyright.Dates))
	}
	if copyright.Dates[0] != "2012" {
		t.Fatalf("result.Copyrights[0].Dates[0] should have been 2012, was %s", copyright.Dates[0])
	}
	if copyright.Name != "A Pinpoint PBC" {
		t.Fatalf("result.Copyrights[0].Name should have been \"A Pinpoint PBC\", was %s", copyright.Name)
	}
}

func TestLineByLineReader(t *testing.T) {
	config := dialect.CreateDefaultConfiguration()
	ex, err := dialect.CreateLineByLineExaminer("JavaScript", "foo.js", config)
	if err != nil {
		t.Fatal(err)
	}
	// pass true as 2nd argument to indicate the EOF
	line, err := ex.ProcessLine([]byte("var a = 1"), true)
	if err != nil {
		t.Fatal(err)
	}
	if line == nil {
		t.Fatal("line was nil")
	}
	if line.LineNumber != 1 {
		t.Fatalf("expected LineNumber to be 1, was %d", line.LineNumber)
	}
	if line.EOF == false {
		t.Fatal("expected EOF to be true, was false")
	}
	if ex.Result == nil {
		t.Fatal("Result was nil")
	}
	if ex.Result.Blanks != 0 {
		t.Fatalf("Blanks should have been 0, was %d", ex.Result.Blanks)
	}
	if ex.Result.Loc != 1 {
		t.Fatalf("Loc should have been 1, was %d", ex.Result.Loc)
	}
	if ex.Result.Sloc != 1 {
		t.Fatalf("Sloc should have been 1, was %d", ex.Result.Sloc)
	}
	if ex.Result.Comments != 0 {
		t.Fatalf("Comments should have been 0, was %d", ex.Result.Comments)
	}
	if ex.Result.IsTest {
		t.Fatal("IsTest should have been false, was true")
	}
}
