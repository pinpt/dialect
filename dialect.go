package dialect

import (
	"bufio"
	"io"

	"github.com/pinpt/dialect/pkg/types"
)

var (
	Build   string
	Version string
)

// CreateDefaultConfiguration will return a default configuration
func CreateDefaultConfiguration() *types.DialectConfiguration {
	return &types.DialectConfiguration{}
}

// CreateConfigurationWithCallback returns a default configuration with a callback
func CreateConfigurationWithCallback(callback types.DialectResultCallback) *types.DialectConfiguration {
	return &types.DialectConfiguration{
		Callback: callback,
	}
}

// CreateLineByLineExaminer returns an interface which can be called with each line using the ProcessLine function
func CreateLineByLineExaminer(language string, filename string, config *types.DialectConfiguration) (*types.DialectContext, error) {
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
func Examine(language string, filename string, reader io.Reader, config *types.DialectConfiguration) (*types.DialectResult, error) {
	ctx, err := CreateLineByLineExaminer(language, filename, config)
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
func DetectFrameworks(directory string) ([]*types.DialectFramework, error) {
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
