package objc

import (
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/languages/objc/ocunit"
	"github.com/pinpt/dialect/pkg/languages/objc/xctest"
	"github.com/pinpt/dialect/pkg/types"
)

type ObjectiveCExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(filename string, line *types.DialectLine) bool {
	return xctest.IsTest(line) || ocunit.IsTest(line)
}

func (e *ObjectiveCExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(filename, line)
	}
	return nil
}

func (e *ObjectiveCExaminer) NewExaminer() types.DialectExaminer {
	ex := new(ObjectiveCExaminer)
	return ex
}

func init() {
	ex := &ObjectiveCExaminer{}
	types.RegisterExaminer("Objective-C", ex)
	types.RegisterExaminer("Objective-C++", ex)
}
