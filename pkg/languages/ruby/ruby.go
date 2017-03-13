package ruby

import (
	"github.com/pinpt/dialect/pkg/languages/ruby/selenium"
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

type RubyExaminer struct {
	inDoubleComment bool
	done            bool
}

func isTest(line *types.DialectLine) bool {
	return selenium.IsTest(line)
}

func (e *RubyExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	lineBuf := strings.TrimSpace(line.Contents)
	if lineBuf == "__END__" {
		e.done = true
	} else if e.done {
		line.IsComment = true
	} else if e.inDoubleComment {
		line.IsComment = true
		if strings.HasPrefix(lineBuf, "=end") {
			// ending of a double line comment
			e.inDoubleComment = false
		}
	} else if strings.HasPrefix(lineBuf, "#") {
		// a single line comment
		line.IsComment = true
	} else if strings.HasPrefix(lineBuf, "=begin") {
		line.IsComment = true
		e.inDoubleComment = true
	} else {
		// this must be code
		line.IsCode = true
	}
	if line.IsTest == false {
		line.IsTest = isTest(line)
	}
	return nil
}

func (e *RubyExaminer) NewExaminer() types.DialectExaminer {
	ex := new(RubyExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Ruby", &RubyExaminer{})
}
