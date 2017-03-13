package nodejs

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

const (
	NAME     = "nodejs"
	FILENAME = "package.json"
	TYPE     = types.DialectFrameworkLanguage
)

var RESULT = []*types.DialectFramework{{Name: NAME, Type: TYPE}}

type NodeJSProcessor struct {
}

func (p *NodeJSProcessor) Detect(directory string) ([]*types.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	types.RegisterFrameworkProcessor(NAME, &NodeJSProcessor{})
}
