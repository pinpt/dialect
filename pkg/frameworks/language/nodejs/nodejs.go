package nodejs

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

const (
	NAME     = "nodejs"
	FILENAME = "package.json"
	TYPE     = dialect.DialectFrameworkLanguage
)

var RESULT = []*dialect.DialectFramework{{Name: NAME, Type: TYPE}}

type NodeJSProcessor struct {
}

func (p *NodeJSProcessor) Detect(directory string) ([]*dialect.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	dialect.RegisterFrameworkProcessor(NAME, &NodeJSProcessor{})
}
