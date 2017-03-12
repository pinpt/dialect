package nodejs

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

const NAME = "nodejs"

type NodeJSProcessor struct {
}

func (p *NodeJSProcessor) Detect(directory string) ([]*dialect.DialectFramework, error) {
	if pkg.FileExists(directory, "package.json") {
		return []*dialect.DialectFramework{{Name: NAME}}, nil
	}
	return nil, nil
}

func init() {
	dialect.RegisterFrameworkProcessor(NAME, &NodeJSProcessor{})
}
