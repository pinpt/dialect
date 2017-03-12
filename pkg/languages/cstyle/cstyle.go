package cstyle

import (
	"github.com/pinpt/dialect"
	"strings"
)

type CStyleExaminer struct {
	inDoubleComment bool
}

func (e *CStyleExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
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

func (e *CStyleExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(CStyleExaminer)
	return ex
}

func init() {
	examiner := &CStyleExaminer{}
	dialect.RegisterExaminer("Apex", examiner)
	dialect.RegisterExaminer("AGS Script", examiner)
	dialect.RegisterExaminer("Arduino", examiner)
	dialect.RegisterExaminer("C", examiner)
	dialect.RegisterExaminer("C++", examiner)
	dialect.RegisterExaminer("ChucK", examiner)
	dialect.RegisterExaminer("Cuda", examiner)
	dialect.RegisterExaminer("DTrace", examiner)
	dialect.RegisterExaminer("EQ", examiner)
	dialect.RegisterExaminer("Game Maker Language", examiner)
	dialect.RegisterExaminer("Kotlin", examiner)
	dialect.RegisterExaminer("Metal", examiner)
	dialect.RegisterExaminer("OpenCL", examiner)
	dialect.RegisterExaminer("Scala", examiner)
	dialect.RegisterExaminer("Squirrel", examiner)
	dialect.RegisterExaminer("Unified Parallel C", examiner)
	dialect.RegisterExaminer("Uno", examiner)
	dialect.RegisterExaminer("UnrealScript", examiner)
	dialect.RegisterExaminer("XC", examiner)
	dialect.RegisterExaminer("XS", examiner)
}
