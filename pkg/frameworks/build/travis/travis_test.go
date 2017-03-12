package travis

import (
	"testing"

	"github.com/pinpt/dialect"
)

func TestFrameworkTravis(t *testing.T) {
	frameworks, err := dialect.DetectFrameworks("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	if len(frameworks) != 1 {
		t.Fatalf("expected 1 frameworks but found %d", len(frameworks))
	}
	if frameworks[0].Name != NAME {
		t.Fatalf("expected %s framework name", NAME)
	}
	if frameworks[0].Type != TYPE {
		t.Fatalf("expected %s framework type", TYPE)
	}
}
