package json

import (
	"github.com/pinpt/dialect/pkg/types"
)

type JSONExaminer struct {
}

func (e *JSONExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	// all lines are code in JSON (excluding blank lines which is handled by the caller)
	line.IsCode = true
	return nil
}

func (e *JSONExaminer) NewExaminer() types.DialectExaminer {
	ex := new(JSONExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("JSON", &JSONExaminer{})
}
