package xml

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

type XMLExaminer struct {
	inDoubleComment bool
}

func (e *XMLExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	lineBuf := strings.TrimSpace(line.Contents)
	if e.inDoubleComment {
		line.IsComment = true
		if strings.HasPrefix(lineBuf, "-->") {
			// ending of a double line comment
			e.inDoubleComment = false
		}
	} else if strings.HasPrefix(lineBuf, "<!--") && strings.HasSuffix(lineBuf, " -->") {
		// a single line comment
		line.IsComment = true
	} else if strings.HasPrefix(lineBuf, "<!--") {
		line.IsComment = true
		e.inDoubleComment = true
	} else {
		// this must be code
		line.IsCode = true
	}
	return nil
}

func (e *XMLExaminer) NewExaminer() types.DialectExaminer {
	ex := new(XMLExaminer)
	return ex
}

func init() {
	ex := &XMLExaminer{}
	types.RegisterExaminer("XML", ex)
	types.RegisterExaminer("XHTML", ex)
	types.RegisterExaminer("HTML", ex)
}
