package ava

import (
	"github.com/pinpt/dialect"
	"regexp"
)

var (
	containsAva  = regexp.MustCompile("['\"]ava['\"]")
	containsTest = regexp.MustCompile("test\\s*\\(")
)

func IsTest(line *dialect.DialectLine) bool {
	return containsAva.MatchString(line.Buffer) && containsTest.MatchString(line.Buffer)
}
