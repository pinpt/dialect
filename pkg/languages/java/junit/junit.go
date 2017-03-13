package junit

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

func IsTest(line *types.DialectLine) bool {
	return strings.Contains(line.Contents, "import junit.framework.") ||
		strings.Contains(line.Contents, "import org.junit.")
}
