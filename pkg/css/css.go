package css

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg/cstyle"
)

type CSSExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func (e *CSSExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	return e.Delegate.Examine(language, filename, line)
}

func (e *CSSExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(CSSExaminer)
	return ex
}

func init() {
	ex := &CSSExaminer{}
	dialect.RegisterExaminer("CSS", ex)
	dialect.RegisterExaminer("SASS", ex)
}
