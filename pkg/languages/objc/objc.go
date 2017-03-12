package objc

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/languages/objc/ocunit"
	"github.com/pinpt/dialect/pkg/languages/objc/xctest"
)

type ObjectiveCExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(filename string, line *dialect.DialectLine) bool {
	return xctest.IsTest(line) || ocunit.IsTest(line)
}

func (e *ObjectiveCExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(filename, line)
	}
	return nil
}

func (e *ObjectiveCExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(ObjectiveCExaminer)
	return ex
}

func init() {
	ex := &ObjectiveCExaminer{}
	dialect.RegisterExaminer("Objective-C", ex)
	dialect.RegisterExaminer("Objective-C++", ex)
}
