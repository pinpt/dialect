package vb

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type VBScriptExaminer struct {
}

func (e *VBScriptExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.SingleSymbolProcessor("! ", line)
	return nil
}

func (e *VBScriptExaminer) NewExaminer() types.DialectExaminer {
	ex := new(VBScriptExaminer)
	return ex
}

func init() {
	ex := &VBScriptExaminer{}
	types.RegisterExaminer("VBScript", ex)
	types.RegisterExaminer("Visual Basic", ex)
}
