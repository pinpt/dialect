package r

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

type RExaminer struct {
	inDoubleComment bool
}

func (e *RExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	lineBuf := strings.TrimSpace(line.Contents)
	if e.inDoubleComment {
		// ending of a double comment
		if strings.HasSuffix(lineBuf, "\"") {
			e.inDoubleComment = false
		}
		line.IsComment = true
	} else {
		if strings.HasPrefix(lineBuf, "# ") && e.inDoubleComment == false {
			// a single line comment
			line.IsComment = true
		} else if strings.HasPrefix(lineBuf, "\"") {
			if strings.HasSuffix(lineBuf, "\"") == false {
				// beginning of a double line comment
				e.inDoubleComment = true
			}
			line.IsComment = true
		} else {
			// regular code
			line.IsCode = true
		}
	}
	return nil
}

func (e *RExaminer) NewExaminer() types.DialectExaminer {
	ex := new(RExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("R", &RExaminer{})
}
