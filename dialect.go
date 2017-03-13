package dialect

import (
	"bufio"
	"io"

	_ "github.com/pinpt/dialect/pkg/frameworks"
	"github.com/pinpt/dialect/pkg/implementation"
	"github.com/pinpt/dialect/pkg/types"
)

var (
	Build   string
	Version string
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
	types.RegisterDialectImplementation(implementation.New())
}
