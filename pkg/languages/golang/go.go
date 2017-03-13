package golang

import (
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

type GoExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func isTest(filename string, line *types.DialectLine) bool {
	return strings.HasSuffix(filename, "_test.go")
}

func (e *GoExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	// handle the first line build flag as regular code
	if line.LineNumber == 1 && strings.HasPrefix(line.Contents, "// +build ") {
		line.IsCode = true
		return nil
	}
	// handle any sys flags or go build instructions as regular code
	if strings.HasPrefix(line.Contents, "//sys ") || strings.HasPrefix(line.Contents, "//go:") {
		line.IsCode = true
		return nil
	}
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	if line.IsTest == false {
		line.IsTest = isTest(filename, line)
	}
	return nil
}

func (e *GoExaminer) NewExaminer() types.DialectExaminer {
	ex := new(GoExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Go", &GoExaminer{})
}
