package haskell

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type HaskellExaminer struct {
}

func (e *HaskellExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.SingleSymbolProcessor("-- ", line)
	return nil
}

func (e *HaskellExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(HaskellExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("Haskell", &HaskellExaminer{})
}
