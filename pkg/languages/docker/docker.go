package docker

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type DockerFileExaminer struct {
}

func (e *DockerFileExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.SingleSymbolProcessor("#", line)
	return nil
}

func (e *DockerFileExaminer) NewExaminer() types.DialectExaminer {
	ex := new(DockerFileExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Dockerfile", &DockerFileExaminer{})
}
