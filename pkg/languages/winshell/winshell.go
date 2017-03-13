package winshell

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type WinShellExaminer struct {
}

func (e *WinShellExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.SingleSymbolProcessor("REM ", line)
	return nil
}

func (e *WinShellExaminer) NewExaminer() types.DialectExaminer {
	ex := new(WinShellExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("WinShell", &WinShellExaminer{})
}
