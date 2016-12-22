package php

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg/cstyle"
	"github.com/pinpt/dialect/pkg/php/selenium"
)

type PHPExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(line *dialect.DialectLine) bool {
	return selenium.IsTest(line)
}

func (e *PHPExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *PHPExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(PHPExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("PHP", &PHPExaminer{})
}
