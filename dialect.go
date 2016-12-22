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

// DialectResultCallback is a callback function for receiving per line results as the code is being examined
type DialectResultCallback func(language string, line *DialectLine) error

// DialectExaminer is the interface that language processors must implement to handle details about a specific language
type DialectExaminer interface {
	NewExaminer() DialectExaminer
	Examine(language string, filename string, line *DialectLine) error
}

var examiners = make(map[string]DialectExaminer)

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

// Examine function is used to detect information about the source code
func Examine(language string, filename string, reader io.Reader, config *DialectConfiguration) (*DialectResult, error) {
	ex := examiners[language]
	if ex == nil {
		ex = examiners["*"]
	}
	if ex == nil {
		return nil, errors.New("the default dialect wasn't registered")
	}
	var LineNumber int
	result := &DialectResult{}
	eol := byte('\n')
	buf := bufio.NewReader(reader)
	var b bytes.Buffer
	for {
		LineNumber++
		contents, err := buf.ReadString('\n')
		eof := err == io.EOF
		b.WriteString(contents)
		b.WriteByte(eol)
		line := &DialectLine{
			EOF:        eof,
			LineNumber: LineNumber,
			Contents:   contents,
			IsBlank:    strings.TrimSpace(contents) == "",
			Buffer:     b.String(),
			Config:     config,
		}
		if line.IsBlank == false {
			if err := ex.Examine(language, filename, line); err != nil {
				return nil, err
			}
		}
		if config != nil && config.Callback != nil {
			if err := config.Callback(language, line); err != nil {
				return nil, err
			}
		}
		if line.IsBlank {
			result.Blanks++
		}
		if line.IsComment {
			result.Comments++
			if config != nil && config.DetectCopyrights {
				comment, err := copyright.ParseCopyright(line.Contents)
				if err != nil {
					return nil, err
				}
				if comment != nil && comment.Found {
					if result.Copyrights == nil {
						result.Copyrights = make([]*copyright.CopyrightResult, 0)
					}
					result.Copyrights = append(result.Copyrights, comment)
				}
			}
		}
		if line.IsCode {
			result.Sloc++
		}
		if line.IsTest {
			result.IsTest = true
		}
		result.Loc++
		if eof {
			break
		}
	}
	return result, nil
}

// RegisterExaminer function is used to register an implementation of the DialectExaminer interface
func RegisterExaminer(language string, examiner DialectExaminer) {
	examiners[language] = examiner
	// add lowercase if not the same
	lc := strings.ToLower(language)
	if lc != language {
		examiners[lc] = examiner
	}
}
