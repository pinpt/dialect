package php

import (
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/languages/php/selenium"
	"github.com/pinpt/dialect/pkg/types"
)

type PHPExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(line *types.DialectLine) bool {
	return selenium.IsTest(line)
}

func (e *PHPExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *PHPExaminer) NewExaminer() types.DialectExaminer {
	ex := new(PHPExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("PHP", &PHPExaminer{})
}
