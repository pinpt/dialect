package sql

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg"
)

type SQLExaminer struct {
}

func (e *SQLExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	pkg.SingleSymbolProcessor("-- ", line)
	return nil
}

func (e *SQLExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(SQLExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("SQL", &SQLExaminer{})
}
