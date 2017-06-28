# Dialect [![Build Status](https://travis-ci.org/pinpt/dialect.svg?branch=master)](https://travis-ci.org/pinpt/dialect)

Dialect is a framework for processing source code to determine basis linguistic details such as number of lines of comments, etc.

Dialect implements a language specific set of rules to determine these details.

If Dialect cannot determine the language or doesn't have a specific implementation, it will simply return the number of lines only.

## Installation

	go get github.com/pinpt/dialect

## Usage

The most simple usage is to invoke `Examine` passing it the language, filename and a `io.Reader` to the source code.

```go
package main

import (
	"fmt"
	"github.com/pinpt/dialect"
	"strings"
)

func main() {
	reader := strings.NewReader("var a = 1")
	config := dialect.CreateDefaultConfiguration()
	result, err := dialect.Examine("JavaScript", "test.js", reader, config)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
```

The result is a struct `DialectResult` which provides the details of what the examiner found.

```go
type DialectResult struct {
	Blanks     int
	Comments   int
	Loc        int
	Sloc       int
	IsTest     bool
	Copyrights []*copyright.CopyrightResult
}
```

The fields are the following:

- _Blanks_ the number of blank lines detected
- _Comments_ the number of comment lines detected
- _Loc_ the total number of lines detected (_Blanks_ + _Comments_ + _Sloc_ = _Loc_)
- _Sloc_ the number of source lines detected (_Sloc_ = _Loc_ - _Comments_ - _Blanks_)
- _IsTest_ a boolean to indicate if the examiner detected the file to be a test file
- _Copyrights_ an array of copyrights found in the file (only available if `DetectCopyrights` is set in config)

## Detecting line by line

Instead of using a Reader, you can feed data line by line.

```go
config := dialect.CreateDefaultConfiguration()
ex, err := dialect.CreateLineByLineExaminer("JavaScript", "foo.js", config)
if err != nil {
	panic(err)
}
// pass true as 2nd argument to indicate the EOF
_, err = ex.ProcessLine("var a = 1", true)
if err != nil {
	panic(err)
}
// results will be in ex.Results
```

## Detecting Copyrights

You can detect Copyrights in source code by setting the configuration field `DetectCopyrights` to `true`.

```go
config := dialect.CreateDefaultConfiguration()
config.DetectCopyrights = true
```

If any copyrights are detected in the source code, the `Copyrights` field will be an array of `*CopyrightResult`.  If none are found, the `Copyrights` will be `nil`.

## Implementing a missing language

To implement your own language or override an existing language, you can use the function `RegisterExaminer` and implement
the interface `DialectExaminer`. See any of the existing implementations for a good example.

If you do implement a new language or fix an existing one, please consider sending us a Pull Request so that we can merge it!

## License

Licensed under the MIT License. Copyright (c) 2016-2017 PinPT, Inc.
