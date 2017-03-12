package docker

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type DockerFileExaminer struct {
}

func (e *DockerFileExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.SingleSymbolProcessor("#", line)
	return nil
}

func (e *DockerFileExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(DockerFileExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("Dockerfile", &DockerFileExaminer{})
}
