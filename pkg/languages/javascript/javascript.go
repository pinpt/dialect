package javascript

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/languages/javascript/ava"
	"github.com/pinpt/dialect/pkg/languages/javascript/cucumber"
	"github.com/pinpt/dialect/pkg/languages/javascript/mocha"
	"github.com/pinpt/dialect/pkg/languages/javascript/selenium"
)

type JavaScriptExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(line *dialect.DialectLine) bool {
	return mocha.IsTest(line) || selenium.IsTest(line) || cucumber.IsTest(line) || ava.IsTest(line)
}

func (e *JavaScriptExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *JavaScriptExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(JavaScriptExaminer)
	return ex
}

func init() {
	ex := &JavaScriptExaminer{}
	dialect.RegisterExaminer("JavaScript", ex)
	dialect.RegisterExaminer("ECMAScript", ex)
}
