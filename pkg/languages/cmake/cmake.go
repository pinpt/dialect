package cmake

import (
	"github.com/pinpt/dialect"
	"strings"
)

type CMakeExaminer struct {
	inDoubleComment bool
}

func (e *CMakeExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	lineBuf := strings.TrimSpace(line.Contents)
	if strings.HasPrefix(lineBuf, "# ") {
		// a single line comment
		line.IsComment = true
	} else {
		// regular code
		line.IsCode = true
	}

	return nil
}

func (e *CMakeExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(CMakeExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("CMake", &CMakeExaminer{})
}
