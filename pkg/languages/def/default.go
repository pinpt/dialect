package def

import (
	"github.com/pinpt/dialect/pkg/types"
)

type DefaultExaminer struct {
}

func (e *DefaultExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	return nil
}

func (e *DefaultExaminer) NewExaminer() types.DialectExaminer {
	ex := new(DefaultExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("*", &DefaultExaminer{})
}
