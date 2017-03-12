package circleci

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

const (
	NAME     = "circleci"
	FILENAME = "circle.yml"
	TYPE     = dialect.DialectFrameworkBuild
)

var RESULT = []*dialect.DialectFramework{{Name: NAME, Type: TYPE}}

type CircleCIProcessor struct {
}

func (p *CircleCIProcessor) Detect(directory string) ([]*dialect.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	dialect.RegisterFrameworkProcessor(NAME, &CircleCIProcessor{})
}
