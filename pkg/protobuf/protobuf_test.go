package protobuf

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/protobuf"
)

func TestProtobuf(t *testing.T) {
	reader := strings.NewReader(`syntax = "proto3";

// comment
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
`)
	result, err := dialect.Examine("Protocol Buffer", "foo.proto", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 9 {
		t.Fatalf("result.Loc should have been 9, was %d", result.Loc)
	}
	if result.Sloc != 6 {
		t.Fatalf("result.Sloc should have been 6, was %d", result.Sloc)
	}
	if result.Comments != 1 {
		t.Fatalf("result.Comments should have been 1, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
