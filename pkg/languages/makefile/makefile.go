package makefile

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type MakefileExaminer struct {
}

func (e *MakefileExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.SingleSymbolProcessor("#", line)
	return nil
}

func (e *MakefileExaminer) NewExaminer() types.DialectExaminer {
	ex := new(MakefileExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Makefile", &MakefileExaminer{})
}
