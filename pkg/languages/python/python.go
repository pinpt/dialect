package python

import (
	"github.com/pinpt/dialect/pkg/languages/python/selenium"
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

type PythonExaminer struct {
	inDoubleComment bool
}

func isTest(line *types.DialectLine) bool {
	return selenium.IsTest(line)
}

func (e *PythonExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	lineBuf := strings.TrimSpace(line.Contents)
	if e.inDoubleComment {
		// ending of a double comment
		if strings.HasPrefix(lineBuf, "\"\"\"") {
			e.inDoubleComment = false
		}
		line.IsComment = true
	} else {
		if strings.HasPrefix(lineBuf, "# ") {
			// a single line comment
			line.IsComment = true
		} else if strings.HasPrefix(lineBuf, "\"\"\"") {
			// beginning of a double line comment
			e.inDoubleComment = true
			line.IsComment = true
		} else {
			// regular code
			line.IsCode = true
		}
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *PythonExaminer) NewExaminer() types.DialectExaminer {
	ex := new(PythonExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Python", &PythonExaminer{})
}
