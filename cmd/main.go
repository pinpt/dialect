package main

import (
	"fmt"
	"os"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/languages"
)

func dump(result *dialect.DialectResult, fn string) {
	fmt.Printf("Statistics for %s:\n\n", fn)
	fmt.Printf("     SLOC:  %5d\n", result.Sloc)
	fmt.Printf(" Comments:  %5d\n", result.Comments)
	fmt.Printf("   Blanks:  %5d\n", result.Blanks)
	fmt.Println("           ------")
	fmt.Printf("      LOC:  %5d\n", result.Loc)
	fmt.Printf("     Test:  %5v\n", result.IsTest)
	fmt.Println()
}

func sum(a *dialect.DialectResult, b *dialect.DialectResult) {
	a.Sloc += b.Sloc
	a.Loc += b.Loc
	a.Comments += b.Comments
	a.Blanks += b.Blanks
	if b.IsTest {
		a.IsTest = true
	}
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: dialect <language> <filename>")
		os.Exit(1)
	}

	language := os.Args[1]
	fn := os.Args[2]

	stats, err := os.Stat(fn)

	if os.IsNotExist(err) {
		fmt.Printf("The file \"%s\" doesn't exist\n", fn)
		os.Exit(1)
	}

	if stats.IsDir() {
		//FIXME
	} else {
		reader, err := os.Open(fn)
		if err != nil {
			fmt.Printf("error reading %s. %v\n", fn, err)
			os.Exit(1)
		}
		result, err := dialect.Examine(language, fn, reader, nil)
		if err != nil {
			fmt.Printf("error examining %s. %v\n", fn, err)
			os.Exit(1)
		}
		dump(result, fn)
	}
}
