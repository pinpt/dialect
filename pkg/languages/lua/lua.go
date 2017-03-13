package lua

import (
	"github.com/pinpt/dialect/pkg/languages/lua/busted"
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

type LuaExaminer struct {
	inDoubleComment bool
}

func isTest(line *types.DialectLine) bool {
	return busted.IsTest(line)
}

func (e *LuaExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	lineBuf := strings.TrimSpace(line.Contents)
	if e.inDoubleComment {
		// ending of a double comment
		if strings.HasPrefix(lineBuf, "--]") {
			e.inDoubleComment = false
		}
		line.IsComment = true
	} else {
		if strings.HasPrefix(lineBuf, "--") {
			// a single line comment
			line.IsComment = true
			if strings.HasPrefix(lineBuf, "--[") {
				// beginning of a double line comment
				e.inDoubleComment = true
			}
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

func (e *LuaExaminer) NewExaminer() types.DialectExaminer {
	ex := new(LuaExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Lua", &LuaExaminer{})
}
