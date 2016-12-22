package wscript

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/wscript"
)

func TestWScriptSingleLine(t *testing.T) {
	reader := strings.NewReader("Set WshNetwork = WScript.CreateObject(\"WScript.Network\")")
	result, err := dialect.Examine("WScript", "foo.wsh", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 1 {
		t.Fatalf("result.Loc should have been 1, was %d", result.Loc)
	}
	if result.Sloc != 1 {
		t.Fatalf("result.Sloc should have been 1, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 0 {
		t.Fatalf("result.Blanks should have been 0, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestWScriptMultiLine(t *testing.T) {
	reader := strings.NewReader(`On Error Resume Next
Set WshNetwork = WScript.CreateObject("WScript.Network")
WshNetwork.MapNetworkDrive "Z:", "\\RemoteServer\Public" '* Map drive Z.
If Err.Number <> 0 Then '* Check to make sure the operation succeeded.
    Err.Clear
    Wscript.Echo "The drive could not be mapped."
End If

'* Map drive Z.
WshNetwork.MapNetworkDrive "Z:", "\\RemoteServer\Public"

'* Check to make sure the operation succeeded.
If Err.Number <> 0 Then
    Err.Clear
    Wscript.Echo "The drive could not be mapped."
End If

REM this is a comment`)
	result, err := dialect.Examine("WScript", "foo.wsh", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 18 {
		t.Fatalf("result.Loc should have been 1, was %d", result.Loc)
	}
	if result.Sloc != 12 {
		t.Fatalf("result.Sloc should have been 12, was %d", result.Sloc)
	}
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 3, was %d", result.Comments)
	}
	if result.Blanks != 3 {
		t.Fatalf("result.Blanks should have been 3, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

/*
 */
