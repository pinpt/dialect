package ocunit

import (
	"github.com/pinpt/dialect"
	"strings"
)

func IsTest(line *dialect.DialectLine) bool {
	return strings.Contains(line.Buffer, "SenTestingKit") ||
		strings.Contains(line.Buffer, "SenTestCase") ||
		strings.Contains(line.Buffer, "STAssertTrue") ||
		strings.Contains(line.Buffer, "STAssertFalse") ||
		strings.Contains(line.Buffer, "STAssertEquals") ||
		strings.Contains(line.Buffer, "STAssertNil") ||
		strings.Contains(line.Buffer, "STAssertNotNil")
}
