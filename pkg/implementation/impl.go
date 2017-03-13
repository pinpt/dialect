package implementation

import (
	"bufio"
	"io"

	"github.com/pinpt/dialect/pkg/types"
)

type DialectImplementation struct {
}

// CreateLineByLineExaminer returns an interface which can be called with each line using the ProcessLine function
func (i *DialectImplementation) CreateLineByLineExaminer(language string, filename string, config *types.DialectConfiguration) (*types.DialectContext, error) {
	ex, err := types.ExaminerForLanguage(language)
	if err != nil {
		return nil, err
	}
	ctx := &types.DialectContext{
		Language: language,
		Filename: filename,
		Config:   config,
		Examiner: ex,
		Result:   &types.DialectResult{},
	}
	return ctx, nil
}

// Examine is used to detect information about the source code
func (i *DialectImplementation) Examine(language string, filename string, reader io.Reader, config *types.DialectConfiguration) (*types.DialectResult, error) {
	ctx, err := i.CreateLineByLineExaminer(language, filename, config)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(reader)
	for {
		contents, err := buf.ReadString('\n')
		eof := err == io.EOF
		if err != nil && eof == false {
			return nil, err
		}
		_, err = ctx.ProcessLine([]byte(contents), eof)
		if err != nil {
			return nil, err
		}
		if eof {
			break
		}
	}
	return ctx.Result, nil
}

// Detect will attempt to detect all the frameworks for a given project source directory
func (i *DialectImplementation) DetectFrameworks(directory string) ([]*types.DialectFramework, error) {
	frameworks := make([]*types.DialectFramework, 0)
	for _, processor := range types.Processors() {
		results, err := processor.Detect(directory)
		if err != nil {
			return nil, err
		}
		for _, r := range results {
			frameworks = append(frameworks, r)
		}
	}
	if len(frameworks) > 0 {
		return frameworks, nil
	}
	return nil, nil
}

// New returns a new implementation of Dialect
func New() types.Dialect {
	return &DialectImplementation{}
}
