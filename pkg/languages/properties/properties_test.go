package properties

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect/pkg/implementation"
	dialect "github.com/pinpt/dialect/pkg/types"
)

func init() {
	dialect.RegisterDialectImplementation(implementation.New())
}

func TestProperties(t *testing.T) {
	reader := strings.NewReader(`
# You are reading the ".properties" entry.
! The exclamation mark can also mark text as comments.
# The key characters =, and : should be written with
# a preceding backslash to ensure that they are properly loaded.
# However, there is no need to preceede the value characters =, and : by a backslash.
website = https://en.wikipedia.org/
language = English
# The backslash below tells the application to continue reading
# the value onto the next line.
message = Welcome to \
          Wikipedia!
# Add spaces to the key
key\ with\ spaces = This is the value that could be looked up with the key "key with spaces".
# Unicode
tab : \u0009
# If you want your property to include a backslash, it should be escaped by another backslash
path=c:\\wiki\\templates
# However, some editors will handle this automatically
`)
	result, err := dialect.Examine("Properties", "foo.properties", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 20 {
		t.Fatalf("result.Loc should have been 20, was %d", result.Loc)
	}
	if result.Sloc != 7 {
		t.Fatalf("result.Sloc should have been 7, was %d", result.Sloc)
	}
	if result.Comments != 11 {
		t.Fatalf("result.Comments should have been 11, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
}
