package selenium

import (
	"github.com/pinpt/dialect"
	"strings"
)

func IsTest(line *dialect.DialectLine) bool {
	return strings.Contains(line.Contents, "Facebook\\WebDriver\\Remote\\RemoteWebDriver")
}
