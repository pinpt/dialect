package jenkins

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

const (
	NAME     = "jenkins"
	FILENAME = "Jenkinsfile"
	TYPE     = dialect.DialectFrameworkBuild
)

var RESULT = []*dialect.DialectFramework{{Name: NAME, Type: TYPE}}

type JenkinsProcessor struct {
}

func (p *JenkinsProcessor) Detect(directory string) ([]*dialect.DialectFramework, error) {
	if pkg.FileExists(directory, FILENAME) {
		return RESULT, nil
	}
	return nil, nil
}

func init() {
	dialect.RegisterFrameworkProcessor(NAME, &JenkinsProcessor{})
}
