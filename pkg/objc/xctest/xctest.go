package xctest

import (
	"github.com/pinpt/dialect"
	"strings"
)

func IsTest(line *dialect.DialectLine) bool {
	return strings.Contains(line.Buffer, "XCTestCase") ||
		strings.Contains(line.Buffer, "XCTest.h")
}
