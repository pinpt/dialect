package selenium

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

func IsTest(line *types.DialectLine) bool {
	return strings.Contains(line.Contents, "Facebook\\WebDriver\\Remote\\RemoteWebDriver")
}
