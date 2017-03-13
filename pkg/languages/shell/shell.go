package shell

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type ShellExaminer struct {
}

func (e *ShellExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.SingleSymbolProcessor("#", line)
	return nil
}

func (e *ShellExaminer) NewExaminer() types.DialectExaminer {
	ex := new(ShellExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Shell", &ShellExaminer{})
}
