package wscript

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type WScriptExaminer struct {
}

func (e *WScriptExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.MultiSymbolProcessor([]string{"'", "REM "}, line)
	return nil
}

func (e *WScriptExaminer) NewExaminer() types.DialectExaminer {
	ex := new(WScriptExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("WScript", &WScriptExaminer{})
}
