package pkg

import (
	"github.com/pinpt/dialect"
	"strings"
)

// SingleSymbolProcessor can be used when you only have a single symbol which designates the line is a comment
func SingleSymbolProcessor(commentSymbol string, line *dialect.DialectLine) {
	if strings.HasPrefix(line.Contents, commentSymbol) {
		line.IsComment = true
	} else {
		line.IsCode = true
	}
}

// MultiSymbolProcessor can be used when you have more than one symbol which designates the line is a comment
func MultiSymbolProcessor(commentSymbols []string, line *dialect.DialectLine) {
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
