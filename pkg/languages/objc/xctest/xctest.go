package xctest

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

func IsTest(line *types.DialectLine) bool {
	return strings.Contains(line.Buffer, "XCTestCase") ||
		strings.Contains(line.Buffer, "XCTest.h")
}
