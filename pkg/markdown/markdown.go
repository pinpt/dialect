package markdown

import (
	"github.com/pinpt/dialect"
)

type MarkdownExaminer struct {
}

func (e *MarkdownExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	// all lines are code in JSON (excluding blank lines which is handled by the caller)
	line.IsCode = true
	return nil
}

func (e *MarkdownExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(MarkdownExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("Markdown", &MarkdownExaminer{})
}
