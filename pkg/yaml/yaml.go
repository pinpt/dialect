package yaml

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type YAMLExaminer struct {
}

func (e *YAMLExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.SingleSymbolProcessor("#", line)
	return nil
}

func (e *YAMLExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(YAMLExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("YAML", &YAMLExaminer{})
}
