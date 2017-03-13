package javascript

import (
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/languages/javascript/ava"
	"github.com/pinpt/dialect/pkg/languages/javascript/cucumber"
	"github.com/pinpt/dialect/pkg/languages/javascript/mocha"
	"github.com/pinpt/dialect/pkg/languages/javascript/selenium"
	"github.com/pinpt/dialect/pkg/types"
)

type JavaScriptExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(line *types.DialectLine) bool {
	return mocha.IsTest(line) || selenium.IsTest(line) || cucumber.IsTest(line) || ava.IsTest(line)
}

func (e *JavaScriptExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *JavaScriptExaminer) NewExaminer() types.DialectExaminer {
	ex := new(JavaScriptExaminer)
	return ex
}

func init() {
	ex := &JavaScriptExaminer{}
	types.RegisterExaminer("JavaScript", ex)
	types.RegisterExaminer("ECMAScript", ex)
}
