package yaml

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type YAMLExaminer struct {
}

func (e *YAMLExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.SingleSymbolProcessor("#", line)
	return nil
}

func (e *YAMLExaminer) NewExaminer() types.DialectExaminer {
	ex := new(YAMLExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("YAML", &YAMLExaminer{})
}
