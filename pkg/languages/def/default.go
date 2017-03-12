package def

import (
	"github.com/pinpt/dialect"
)

type DefaultExaminer struct {
}

func (e *DefaultExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	return nil
}

func (e *DefaultExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(DefaultExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("*", &DefaultExaminer{})
}
