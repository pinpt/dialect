package travis

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

const (
	NAME     = "travis"
	FILENAME = ".travis.yml"
	TYPE     = dialect.DialectFrameworkBuild
)

var RESULT = []*dialect.DialectFramework{{Name: NAME, Type: TYPE}}

type TravisProcessor struct {
}

func (p *TravisProcessor) Detect(directory string) ([]*dialect.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	dialect.RegisterFrameworkProcessor(NAME, &TravisProcessor{})
}
