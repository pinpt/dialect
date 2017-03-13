package sql

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type SQLExaminer struct {
}

func (e *SQLExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.SingleSymbolProcessor("-- ", line)
	return nil
}

func (e *SQLExaminer) NewExaminer() types.DialectExaminer {
	ex := new(SQLExaminer)
	return ex
}

func init() {
	ex := &SQLExaminer{}
	types.RegisterExaminer("SQL", ex)
	types.RegisterExaminer("PLSQL", ex)
	types.RegisterExaminer("PLpgSQL", ex)
	types.RegisterExaminer("SQLPL", ex)
}
