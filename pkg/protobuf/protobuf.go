package protobuf

import (
	"github.com/pinpt/dialect"
	"github.com/pinpt/dialect/pkg/cstyle"
)

type ProtoBufExaminer struct {
	Delegate cstyle.CStyleExaminer
}

func (e *ProtoBufExaminer) Examine(language string, filename string, line *dialect.DialectLine) error {
	if err := e.Delegate.Examine(language, filename, line); err != nil {
		return err
	}
	return nil
}

func (e *ProtoBufExaminer) NewExaminer() dialect.DialectExaminer {
	ex := new(ProtoBufExaminer)
	return ex
}

func init() {
	dialect.RegisterExaminer("Protocol Buffer", &ProtoBufExaminer{})
}
