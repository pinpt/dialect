package lua

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/lua"
)

func TestLua(t *testing.T) {
	reader := strings.NewReader(`
		-- hello.lua
		-- the first program in every language

		local myvar= 2
		io.write("Hello world, from ",_VERSION,"!\n")

		--[
		this is a comment lol
		--]
	`)
	result, err := dialect.Examine("Lua", "foo.lua", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 11 {
		t.Fatalf("result.Loc should have been 11, was %d", result.Loc)
	}
	if result.Sloc != 2 {
		t.Fatalf("result.Sloc should have been 2, was %d", result.Sloc)
	}
	if result.Comments != 5 {
		t.Fatalf("result.Comments should have been 5, was %d", result.Comments)
	}
	if result.Blanks != 4 {
		t.Fatalf("result.Blanks should have been 4, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}
