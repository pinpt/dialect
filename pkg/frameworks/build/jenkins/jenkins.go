package jenkins

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

const (
	NAME     = "jenkins"
	FILENAME = "Jenkinsfile"
	TYPE     = types.DialectFrameworkBuild
)

var RESULT = []*types.DialectFramework{{Name: NAME, Type: TYPE}}

type JenkinsProcessor struct {
}

func (p *JenkinsProcessor) Detect(directory string) ([]*types.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	types.RegisterFrameworkProcessor(NAME, &JenkinsProcessor{})
}
