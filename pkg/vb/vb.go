package vb

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type VBScriptExaminer struct {
}

func (e *VBScriptExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.SingleSymbolProcessor("! ", line)
	return nil
}

func (e *VBScriptExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(VBScriptExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("VBScript", &VBScriptExaminer{})
}
