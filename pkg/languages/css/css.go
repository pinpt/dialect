package css

import (
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/types"
)

type CSSExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func (e *CSSExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	return e.Delegate.Examine(language, filename, line)
}

func (e *CSSExaminer) NewExaminer() types.DialectExaminer {
	ex := new(CSSExaminer)
	return ex
}

func init() {
	ex := &CSSExaminer{}
	types.RegisterExaminer("CSS", ex)
	types.RegisterExaminer("Sass", ex)
	types.RegisterExaminer("SCSS", ex)
	types.RegisterExaminer("Stylus", ex)
}
