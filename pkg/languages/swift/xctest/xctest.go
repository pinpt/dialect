package xctest

import (
	"github.com/pinpt/dialect"
	"strings"
)

func IsTest(line *dialect.DialectLine) bool {
	return strings.Contains(line.Contents, "import XCTest") || strings.Contains(line.Buffer, "XCTestCase")
}
