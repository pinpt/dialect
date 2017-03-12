package docker

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

const (
	NAME     = "docker"
	FILENAME = "Dockerfile"
	TYPE     = dialect.DialectFrameworkContainer
)

var RESULT = []*dialect.DialectFramework{{Name: NAME, Type: TYPE}}

type DockerProcessor struct {
}

func (p *DockerProcessor) Detect(directory string) ([]*dialect.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	dialect.RegisterFrameworkProcessor(NAME, &DockerProcessor{})
}
