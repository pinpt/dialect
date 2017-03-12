package makefile

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type MakefileExaminer struct {
}

func (e *MakefileExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.SingleSymbolProcessor("#", line)
	return nil
}

func (e *MakefileExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(MakefileExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("Makefile", &MakefileExaminer{})
}
