package properties

import (
	"github.com/pinpt/dialect/pkg"
	"github.com/pinpt/dialect/pkg/types"
)

type PropertiesExaminer struct {
}

func (e *PropertiesExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	pkg.MultiSymbolProcessor([]string{"#", "!"}, line)
	return nil
}

func (e *PropertiesExaminer) NewExaminer() types.DialectExaminer {
	ex := new(PropertiesExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Properties", &PropertiesExaminer{})
}
