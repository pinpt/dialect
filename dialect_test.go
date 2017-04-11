package dialect

import (
	"bytes"
	"strings"
	"sync"
	"testing"
)

func TestLoadDialect(t *testing.T) {
	config := CreateDefaultConfiguration()
	result, err := Examine("JavaScript", "test.js", strings.NewReader("var a = 1"), config)
	if err != nil {
		t.Fatal(err)
	}
	if result.Blanks != 0 {
		t.Fatalf("expected Blanks to be 0 was %d", result.Blanks)
	}
	if result.Comments != 0 {
		t.Fatalf("expected Comments to be 0 was %d", result.Comments)
	}
	if result.Sloc != 1 {
		t.Fatalf("expected Sloc to be 1 was %d", result.Sloc)
	}
	if result.Loc != 1 {
		t.Fatalf("expected Loc to be 1 was %d", result.Loc)
	}
	if result.IsTest {
		t.Fatalf("expected Test to be false was %v", result.IsTest)
	}
}

func TestConcurrentAccess(t *testing.T) {
	wg := sync.WaitGroup{}
	config := CreateDefaultConfiguration()
	buf := []byte(`/**
* comment
*/
var a = 1;
`)
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := Examine("JavaScript", "test.js", bytes.NewReader(buf), config)
			if err != nil {
				t.Fatal(err)
			}
			if result.Sloc != 1 {
				t.Fatalf("expected Sloc to be 1 but was %d", result.Sloc)
			}
			if result.Loc != 5 {
				t.Fatalf("expected Loc to be 5 but was %d", result.Loc)
			}
			if result.Comments != 3 {
				t.Fatalf("expected Comments to be 3 but was %d", result.Comments)
			}
			if result.Blanks != 1 {
				t.Fatalf("expected Blanks to be 1 but was %d", result.Blanks)
			}
		}()
	}
	wg.Wait()
}
