package xctest

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

func IsTest(line *types.DialectLine) bool {
	return strings.Contains(line.Contents, "import XCTest") || strings.Contains(line.Buffer, "XCTestCase")
}
