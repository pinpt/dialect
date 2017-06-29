package cstyle

import (
	"strings"

	"github.com/pinpt/dialect/pkg/types"
)

type CStyleExaminer struct {
	inDoubleComment bool
}

func (e *CStyleExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	lineBuf := strings.TrimSpace(line.Contents)
	if e.inDoubleComment {
		// ending of a double comment
		if strings.HasSuffix(lineBuf, "*/") {
			e.inDoubleComment = false
		}
		line.IsComment = true
	} else {
		if strings.HasPrefix(lineBuf, "//") {
			// a single line comment
			line.IsComment = true
		} else if strings.HasPrefix(lineBuf, "/*") {
			if strings.HasSuffix(lineBuf, "*/") {
				// inline comment
				line.IsComment = true
			} else {
				// beginning of a double line comment
				e.inDoubleComment = true
				line.IsComment = true
			}
		} else {
			// regular code
			line.IsCode = true
		}
	}
	return nil
}

func (e *CStyleExaminer) NewExaminer() types.DialectExaminer {
	ex := new(CStyleExaminer)
	return ex
}

func init() {
	examiner := &CStyleExaminer{}
	types.RegisterExaminer("Apex", examiner)
	types.RegisterExaminer("AGS Script", examiner)
	types.RegisterExaminer("Arduino", examiner)
	types.RegisterExaminer("C", examiner)
	types.RegisterExaminer("C++", examiner)
	types.RegisterExaminer("ChucK", examiner)
	types.RegisterExaminer("Cuda", examiner)
	types.RegisterExaminer("DTrace", examiner)
	types.RegisterExaminer("Dart", examiner)
	types.RegisterExaminer("EQ", examiner)
	types.RegisterExaminer("Game Maker Language", examiner)
	types.RegisterExaminer("Kotlin", examiner)
	types.RegisterExaminer("Metal", examiner)
	types.RegisterExaminer("OpenCL", examiner)
	types.RegisterExaminer("Scala", examiner)
	types.RegisterExaminer("Squirrel", examiner)
	types.RegisterExaminer("Unified Parallel C", examiner)
	types.RegisterExaminer("Uno", examiner)
	types.RegisterExaminer("UnrealScript", examiner)
	types.RegisterExaminer("XC", examiner)
	types.RegisterExaminer("XS", examiner)
}
