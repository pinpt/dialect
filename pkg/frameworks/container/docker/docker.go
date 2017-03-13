package docker

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

const (
	NAME     = "docker"
	FILENAME = "Dockerfile"
	TYPE     = types.DialectFrameworkContainer
)

var RESULT = []*types.DialectFramework{{Name: NAME, Type: TYPE}}

type DockerProcessor struct {
}

func (p *DockerProcessor) Detect(directory string) ([]*types.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	types.RegisterFrameworkProcessor(NAME, &DockerProcessor{})
}
