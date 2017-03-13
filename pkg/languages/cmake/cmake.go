package cmake

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

type CMakeExaminer struct {
	inDoubleComment bool
}

func (e *CMakeExaminer) Examine(language string, filename string, line *types.DialectLine) error {
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

func (e *CMakeExaminer) NewExaminer() types.DialectExaminer {
	ex := new(CMakeExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("CMake", &CMakeExaminer{})
}
