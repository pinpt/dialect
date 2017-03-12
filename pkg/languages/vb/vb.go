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
	ex := &VBScriptExaminer{}
	dialect.RegisterExaminer("VBScript", ex)
	dialect.RegisterExaminer("Visual Basic", ex)
}
