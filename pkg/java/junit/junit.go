package junit

import (
	"github.com/pinpt/dialect"
	"strings"
)

func IsTest(line *dialect.DialectLine) bool {
	return strings.Contains(line.Contents, "import junit.framework.") ||
		strings.Contains(line.Contents, "import org.junit.")
}
