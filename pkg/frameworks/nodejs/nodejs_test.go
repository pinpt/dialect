package nodejs

import (
	"testing"

	"github.com/pinpt/dialect"
)

func TestNodeJS(t *testing.T) {
	frameworks, err := dialect.DetectFrameworks("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	if len(frameworks) != 1 {
		t.Fatalf("expected 1 frameworks but found %d", len(frameworks))
	}
	if frameworks[0].Name != "nodejs" {
		t.Fatalf("expected nodejs framework")
	}
}
