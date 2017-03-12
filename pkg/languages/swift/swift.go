package swift

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/languages/swift/xctest"
)

type SwiftExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(filename string, line *dialect.DialectLine) bool {
	return xctest.IsTest(line)
}

func (e *SwiftExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(filename, line)
	}
	return nil
}

func (e *SwiftExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(SwiftExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("Swift", &SwiftExaminer{})
}
