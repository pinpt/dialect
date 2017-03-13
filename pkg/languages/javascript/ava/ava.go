package ava

import (
	"github.com/pinpt/dialect/pkg/types"
	"regexp"
)

var (
	containsAva  = regexp.MustCompile("['\"]ava['\"]")
	containsTest = regexp.MustCompile("test\\s*\\(")
)

func IsTest(line *types.DialectLine) bool {
	return containsAva.MatchString(line.Buffer) && containsTest.MatchString(line.Buffer)
}
