package shell

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type ShellExaminer struct {
}

func (e *ShellExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.SingleSymbolProcessor("#", line)
	return nil
}

func (e *ShellExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(ShellExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("Shell", &ShellExaminer{})
}
