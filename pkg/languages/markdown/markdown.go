package markdown

import (
	"github.com/pinpt/dialect/pkg/types"
)

type MarkdownExaminer struct {
}

func (e *MarkdownExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	// all lines are code in JSON (excluding blank lines which is handled by the caller)
	line.IsCode = true
	return nil
}

func (e *MarkdownExaminer) NewExaminer() types.DialectExaminer {
	ex := new(MarkdownExaminer)
	return ex
}

func init() {
	ex := &MarkdownExaminer{}
	types.RegisterExaminer("Markdown", ex)
	types.RegisterExaminer("RMarkdown", ex)
}
