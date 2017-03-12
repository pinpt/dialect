package dialect

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/pinpt/dialect/pkg/copyright"
)

// DialectResult is returned from Examine to describe the code
type DialectResult struct {
	Blanks     int
	Comments   int
	Loc        int
	Sloc       int
	IsTest     bool
	Copyrights []*copyright.CopyrightResult
}

// DialectConfiguration is a struct which is used for specifying configuration
type DialectConfiguration struct {
	Callback         DialectResultCallback
	DetectCopyrights bool
}

func (r *DialectResult) String() string {
	return fmt.Sprintf("result[blanks=%d,comments=%d,sloc=%d,loc=%d,test=%v]", r.Blanks, r.Comments, r.Sloc, r.Loc, r.IsTest)
}

// DialectLine is internally used by the dialect framework to communicate results per line by a DialectExaminer implementation
type DialectLine struct {
	LineNumber int
	IsComment  bool
	IsCode     bool
	IsBlank    bool
	IsTest     bool
	Contents   string
	EOF        bool
	Buffer     string
	Config     *DialectConfiguration
}

// DialectFramework is details about a framework that the project supports
type DialectFramework struct {
	Name string
}

// DialectResultCallback is a callback function for receiving per line results as the code is being examined
type DialectResultCallback func(language string, line *DialectLine) error

// DialectExaminer is the interface that language processors must implement to handle details about a specific language
type DialectExaminer interface {
	NewExaminer() DialectExaminer
	Examine(language string, filename string, line *DialectLine) error
}

// DialectFramework is the interface that framework processors must implement to handle details about a specific framework
type DialectFrameworkProcessor interface {
	Detect(directory string) ([]*DialectFramework, error)
}

type DialectContext struct {
	Language   string
	Filename   string
	Config     *DialectConfiguration
	Examiner   DialectExaminer
	LineNumber int
	Result     *DialectResult
	Buffer     bytes.Buffer
}

const EOL = byte('\n')

func (ctx *DialectContext) ProcessLine(buf []byte, eof bool) (*DialectLine, error) {
	ctx.LineNumber++
	contents := string(buf)
	ctx.Buffer.WriteString(contents)
	ctx.Buffer.WriteByte(EOL)
	line := &DialectLine{
		EOF:        eof,
		LineNumber: ctx.LineNumber,
		Contents:   contents,
		IsBlank:    strings.TrimSpace(contents) == "",
		Buffer:     ctx.Buffer.String(),
		Config:     ctx.Config,
	}
	if line.IsBlank == false {
		if err := ctx.Examiner.Examine(ctx.Language, ctx.Filename, line); err != nil {
			return nil, err
		}
	}
	if ctx.Config != nil && ctx.Config.Callback != nil {
		if err := ctx.Config.Callback(ctx.Language, line); err != nil {
			return nil, err
		}
	}
	if line.IsBlank {
		ctx.Result.Blanks++
	}
	if line.IsComment {
		ctx.Result.Comments++
		if ctx.Config != nil && ctx.Config.DetectCopyrights {
			comment, err := copyright.ParseCopyright(line.Contents)
			if err != nil {
				return nil, err
			}
			if comment != nil && comment.Found {
				if ctx.Result.Copyrights == nil {
					ctx.Result.Copyrights = make([]*copyright.CopyrightResult, 0)
				}
				ctx.Result.Copyrights = append(ctx.Result.Copyrights, comment)
			}
		}
	}
	if line.IsCode {
		ctx.Result.Sloc++
	}
	if line.IsTest {
		ctx.Result.IsTest = true
	}
	ctx.Result.Loc++
	return line, nil
}

var examiners = make(map[string]DialectExaminer)
var processors = make(map[string]DialectFrameworkProcessor)

// CreateDefaultConfiguration will return a default configuration
func CreateDefaultConfiguration() *DialectConfiguration {
	return &DialectConfiguration{}
}

// CreateConfigurationWithCallback returns a default configuration with a callback
func CreateConfigurationWithCallback(callback DialectResultCallback) *DialectConfiguration {
	return &DialectConfiguration{
		Callback: callback,
	}
}

// CreateLineByLineExaminer returns an interface which can be called with each line using the ProcessLine function
func CreateLineByLineExaminer(language string, filename string, config *DialectConfiguration) (*DialectContext, error) {
	ex := examiners[language]
	if ex == nil {
		ex = examiners["*"]
	}
	if ex == nil {
		return nil, errors.New("the default dialect wasn't registered")
	}
	ctx := &DialectContext{
		Language: language,
		Filename: filename,
		Config:   config,
		Examiner: ex,
		Result:   &DialectResult{},
	}
	return ctx, nil
}

// Examine is used to detect information about the source code
func Examine(language string, filename string, reader io.Reader, config *DialectConfiguration) (*DialectResult, error) {
	ctx, err := CreateLineByLineExaminer(language, filename, config)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(reader)
	for {
		contents, err := buf.ReadString('\n')
		eof := err == io.EOF
		if err != nil && eof == false {
			return nil, err
		}
		_, err = ctx.ProcessLine([]byte(contents), eof)
		if err != nil {
			return nil, err
		}
		if eof {
			break
		}
	}
	return ctx.Result, nil
}

// Detect will attempt to detect all the frameworks for a given project source directory
func DetectFrameworks(directory string) ([]*DialectFramework, error) {
	frameworks := make([]*DialectFramework, 0)
	for _, processor := range processors {
		results, err := processor.Detect(directory)
		if err != nil {
			return nil, err
		}
		for _, r := range results {
			frameworks = append(frameworks, r)
		}
	}
	if len(frameworks) > 0 {
		return frameworks, nil
	}
	return nil, nil
}

// RegisterExaminer is used to register an implementation of the DialectExaminer interface
func RegisterExaminer(language string, examiner DialectExaminer) {
	examiners[language] = examiner
	// add lowercase if not the same
	lc := strings.ToLower(language)
	if lc != language {
		examiners[lc] = examiner
	}
}

// RegisterFrameworkProcessor is used to register an implementation of the DialectFrameworkProcessor interface
func RegisterFrameworkProcessor(name string, processor DialectFrameworkProcessor) {
	processors[name] = processor
}
