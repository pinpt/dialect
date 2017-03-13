package pkg

import (
	"os"
	"path"
	"strings"

	"github.com/pinpt/dialect/pkg/types"
)

// SingleSymbolProcessor can be used when you only have a single symbol which designates the line is a comment
func SingleSymbolProcessor(commentSymbol string, line *types.DialectLine) {
	if strings.HasPrefix(line.Contents, commentSymbol) {
		line.IsComment = true
	} else {
		line.IsCode = true
	}
}

// MultiSymbolProcessor can be used when you have more than one symbol which designates the line is a comment
func MultiSymbolProcessor(commentSymbols []string, line *types.DialectLine) {
	var found bool
	for _, commentSymbol := range commentSymbols {
		if strings.HasPrefix(line.Contents, commentSymbol) {
			found = true
			break
		}
	}
	if found {
		line.IsComment = true
	} else {
		line.IsCode = true
	}
}

// FileExists returns true if the filename path exists or false if not
func FileExists(filename ...string) bool {
	fn := path.Join(filename...)
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return false
	}
	return true
}
