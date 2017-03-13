package csharp

import (
	"github.com/pinpt/dialect/pkg/languages/csharp/selenium"
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/types"
)

type CSharpExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(line *types.DialectLine) bool {
	return selenium.IsTest(line)
}

func (e *CSharpExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *CSharpExaminer) NewExaminer() types.DialectExaminer {
	ex := new(CSharpExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("C#", &CSharpExaminer{})
}
