package coffeescript

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

type CoffeeScriptExaminer struct {
	inDoubleComment bool
}

func (e *CoffeeScriptExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	lineBuf := strings.TrimSpace(line.Contents)
	if e.inDoubleComment {
		// ending of a double comment
		if lineBuf == "###" {
			e.inDoubleComment = false
		}
		line.IsComment = true
	} else {
		if lineBuf == "###" {
			// beginning of a double line comment
			e.inDoubleComment = true
			line.IsComment = true
		} else if strings.HasPrefix(lineBuf, "#") {
			// a single line comment
			line.IsComment = true
		} else if e.inDoubleComment == false {
			// regular code
			line.IsCode = true
		}
	}
	return nil
}

func (e *CoffeeScriptExaminer) NewExaminer() types.DialectExaminer {
	ex := new(CoffeeScriptExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("CoffeeScript", &CoffeeScriptExaminer{})
}
