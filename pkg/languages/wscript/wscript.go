package wscript

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type WScriptExaminer struct {
}

func (e *WScriptExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.MultiSymbolProcessor([]string{"'", "REM "}, line)
	return nil
}

func (e *WScriptExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(WScriptExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("WScript", &WScriptExaminer{})
}
