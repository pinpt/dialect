package gopkg

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/gopkg"
)

func TestGoSingleLine(t *testing.T) {
	reader := strings.NewReader(`package foo
var a string

// Foo will do something
func Foo() string {
	return a
}
`)
	result, err := dialect.Examine("Go", "foo.go", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 8 {
		t.Fatalf("result.Loc should have been 8, was %d", result.Loc)
	}
	if result.Sloc != 5 {
		t.Fatalf("result.Sloc should have been 5, was %d", result.Sloc)
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

func TestGoTestCase(t *testing.T) {
	reader := strings.NewReader(`package foo
var a string

func TestFoo(t *testing.T) {
}
`)
	result, err := dialect.Examine("Go", "foo_test.go", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 6 {
		t.Fatalf("result.Loc should have been 6, was %d", result.Loc)
	}
	if result.Sloc != 4 {
		t.Fatalf("result.Sloc should have been 4, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}

func TestGoWithBuildArgs(t *testing.T) {
	reader := strings.NewReader(`// +build linux
package foo
var a string

func TestFoo(t *testing.T) {
}
`)
	result, err := dialect.Examine("Go", "foo_test.go", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 7 {
		t.Fatalf("result.Loc should have been 7, was %d", result.Loc)
	}
	if result.Sloc != 5 {
		t.Fatalf("result.Sloc should have been 5, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}

func TestGoWithSysFlags(t *testing.T) {
	reader := strings.NewReader(`package foo
var a string

//sys cancelIoEx(file syscall.Handle, o *syscall.Overlapped) (err error) = CancelIoEx

func TestFoo(t *testing.T) {
}
`)
	result, err := dialect.Examine("Go", "foo_test.go", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 8 {
		t.Fatalf("result.Loc should have been 8, was %d", result.Loc)
	}
	if result.Sloc != 5 {
		t.Fatalf("result.Sloc should have been 5, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 3 {
		t.Fatalf("result.Blanks should have been 3, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}

func TestGoWithGoFlags(t *testing.T) {
	reader := strings.NewReader(`package foo
var a string

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go file.go

func TestFoo(t *testing.T) {
}
`)
	result, err := dialect.Examine("Go", "foo_test.go", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 8 {
		t.Fatalf("result.Loc should have been 8, was %d", result.Loc)
	}
	if result.Sloc != 5 {
		t.Fatalf("result.Sloc should have been 5, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 3 {
		t.Fatalf("result.Blanks should have been 3, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}
