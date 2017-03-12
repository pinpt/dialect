package java

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/languages/java/junit"
	"github.com/pinpt/dialect/pkg/languages/java/selenium"
)

type JavaExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(line *dialect.DialectLine) bool {
	return junit.IsTest(line) || selenium.IsTest(line)
}

func (e *JavaExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *JavaExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(JavaExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("Java", &JavaExaminer{})
}
