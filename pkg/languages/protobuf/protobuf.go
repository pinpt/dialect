package protobuf

import (
	"github.com/pinpt/dialect/pkg/languages/cstyle"
	"github.com/pinpt/dialect/pkg/types"
)

type ProtoBufExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func (e *ProtoBufExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	return nil
}

func (e *ProtoBufExaminer) NewExaminer() types.DialectExaminer {
	ex := new(ProtoBufExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("Protocol Buffer", &ProtoBufExaminer{})
}
