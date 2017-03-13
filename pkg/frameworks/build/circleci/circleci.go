package circleci

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

const (
	NAME     = "circleci"
	FILENAME = "circle.yml"
	TYPE     = types.DialectFrameworkBuild
)

var RESULT = []*types.DialectFramework{{Name: NAME, Type: TYPE}}

type CircleCIProcessor struct {
}

func (p *CircleCIProcessor) Detect(directory string) ([]*types.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	types.RegisterFrameworkProcessor(NAME, &CircleCIProcessor{})
}
