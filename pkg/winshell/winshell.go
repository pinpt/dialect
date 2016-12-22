package winshell

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type WinShellExaminer struct {
}

func (e *WinShellExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.SingleSymbolProcessor("REM ", line)
	return nil
}

func (e *WinShellExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(WinShellExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("WinShell", &WinShellExaminer{})
}
