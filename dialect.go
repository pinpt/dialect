package dialect

import (
	"io"

	_ "github.com/pinpt/dialect/pkg/frameworks"
	"github.com/pinpt/dialect/pkg/implementation"
	_ "github.com/pinpt/dialect/pkg/languages"
	"github.com/pinpt/dialect/pkg/types"
)

var (
	Build   string
	Version string

	impl types.Dialect
)

// CreateDefaultConfiguration will return a default configuration
func CreateDefaultConfiguration() *types.DialectConfiguration {
	return types.CreateDefaultConfiguration()
}

// CreateConfigurationWithCallback returns a default configuration with a callback
func CreateConfigurationWithCallback(callback types.DialectResultCallback) *types.DialectConfiguration {
	return types.CreateConfigurationWithCallback(callback)
}

func init() {
	impl = implementation.New()
	types.RegisterDialectImplementation(impl)
}

// CreateLineByLineExaminer returns an interface which can be called with each line using the ProcessLine function
func CreateLineByLineExaminer(language string, filename string, config *types.DialectConfiguration) (*types.DialectContext, error) {
	return impl.CreateLineByLineExaminer(language, filename, config)
}

// Examine is used to detect information about the source code
func Examine(language string, filename string, reader io.Reader, config *types.DialectConfiguration) (*types.DialectResult, error) {
	return impl.Examine(language, filename, reader, config)
}

// Detect will attempt to detect all the frameworks for a given project source directory
func DetectFrameworks(directory string) ([]*types.DialectFramework, error) {
	return impl.DetectFrameworks(directory)
}
