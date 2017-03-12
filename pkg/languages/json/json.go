package json

import (
	"github.com/pinpt/dialect"
)

type JSONExaminer struct {
}

func (e *JSONExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	// all lines are code in JSON (excluding blank lines which is handled by the caller)
	line.IsCode = true
	return nil
}

func (e *JSONExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(JSONExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("JSON", &JSONExaminer{})
}
