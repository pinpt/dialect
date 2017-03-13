package ejs

import (
	"github.com/pinpt/dialect/pkg/languages/javascript"
	"github.com/pinpt/dialect/pkg/languages/xml"
	"github.com/pinpt/dialect/pkg/types"
)

type EJSExaminer struct {
	XML xml.XMLExaminer
	JS  javascript.JavaScriptExaminer
}

func (e *EJSExaminer) Examine(language string, filename string, line *types.DialectLine) error {
	// EJS is a combination (usually) of HTML and JS so we're going to use both
	r1 := &types.DialectLine{
		LineNumber: line.LineNumber,
		Contents:   line.Contents,
		Buffer:     line.Buffer,
	}
	r2 := &types.DialectLine{
		LineNumber: line.LineNumber,
		Contents:   line.Contents,
		Buffer:     line.Buffer,
	}
	if err := e.JS.Examine(language, filename, r2); err != nil {
		return err
	}
	if r2.IsCode {
		line.IsCode = true
		return nil
	}
	if r2.IsComment {
		line.IsComment = true
		return nil
	}
	if err := e.XML.Examine(language, filename, r1); err != nil {
		return err
	}
	if r1.IsCode {
		line.IsCode = true
		return nil
	}
	if r1.IsComment {
		line.IsComment = true
		return nil
	}
	return nil
}

func (e *EJSExaminer) NewExaminer() types.DialectExaminer {
	ex := new(EJSExaminer)
	return ex
}

func init() {
	types.RegisterExaminer("EJS", &EJSExaminer{})
}
