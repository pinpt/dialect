package busted

import (
	"github.com/pinpt/dialect/pkg/types"
	"regexp"
)

var (
	containsDescribe = regexp.MustCompile("describe\\s*\\(")
	containsIt       = regexp.MustCompile("it\\s?\\(")
)

func IsTest(line *types.DialectLine) bool {
	return containsDescribe.MatchString(line.Buffer) && containsIt.MatchString(line.Buffer)
}
