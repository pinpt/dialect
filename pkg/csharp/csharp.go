package csharp

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg/csharp/selenium"
	"github.com/pinpt/dialect/pkg/cstyle"
)

type CSharpExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(line *dialect.DialectLine) bool {
	return selenium.IsTest(line)
}

func (e *CSharpExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *CSharpExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(CSharpExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("C#", &CSharpExaminer{})
}
