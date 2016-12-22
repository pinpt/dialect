package mocha

import (
	"github.com/pinpt/dialect"
	"regexp"
)

var (
	containsDescribe = regexp.MustCompile("describe\\s*\\(")
	containsIt       = regexp.MustCompile("it\\s?(\\.only|skip)?\\s?\\(")
)

func IsTest(line *dialect.DialectLine) bool {
	return containsDescribe.MatchString(line.Buffer) && containsIt.MatchString(line.Buffer)
}
