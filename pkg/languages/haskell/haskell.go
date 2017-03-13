package haskell

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type HaskellExaminer struct {
}

func (e *HaskellExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.SingleSymbolProcessor("-- ", line)
	return nil
}

func (e *HaskellExaminer) NewExaminer() types.DialectExaminer {
	ex := new(HaskellExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Haskell", &HaskellExaminer{})
}
