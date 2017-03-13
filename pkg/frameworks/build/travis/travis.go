package travis

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

const (
	NAME     = "travis"
	FILENAME = ".travis.yml"
	TYPE     = types.DialectFrameworkBuild
)

var RESULT = []*types.DialectFramework{{Name: NAME, Type: TYPE}}

type TravisProcessor struct {
}

func (p *TravisProcessor) Detect(directory string) ([]*types.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	types.RegisterFrameworkProcessor(NAME, &TravisProcessor{})
}
