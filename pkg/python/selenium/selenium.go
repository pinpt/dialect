package selenium

import (
	"github.com/pinpt/dialect"
	"strings"
)

func IsTest(line *dialect.DialectLine) bool {
	return strings.Contains(line.Contents, "selenium") && strings.Contains(line.Contents, "webdriver")
}
