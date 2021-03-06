package cucumber

import (
	"github.com/pinpt/dialect/pkg/types"
	"strings"
)

func IsTest(line *types.DialectLine) bool {
	return strings.Contains(line.Buffer, "Cucumber.") &&
		(strings.Contains(line.Buffer, "FeatureParser") ||
			strings.Contains(line.Buffer, "getSupportCodeFns") ||
			strings.Contains(line.Buffer, "FormatterBuilder") ||
			strings.Contains(line.Buffer, "Runtime") || strings.Contains(line.Buffer, "ScenarioFilter"))
}
