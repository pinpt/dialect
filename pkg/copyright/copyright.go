package copyright

import (
	"bytes"
	"regexp"
	"strings"
)

type CopyrightResult struct {
	Found bool
	Dates []string
	Name  string
}

func (c *CopyrightResult) Stringify(symbol string) string {
	if c.Found {
		var b bytes.Buffer
		b.WriteString("Copyright ")
		b.WriteString(symbol)
		b.WriteString(" ")
		if c.Dates != nil && len(c.Dates) > 0 {
			if len(c.Dates) > 1 {
				b.WriteString(strings.Join(c.Dates, "-"))
			} else {
				b.WriteString(c.Dates[0])
			}
			b.WriteString(" ")
		}
		if c.Name != "" {
			b.WriteString("by ")
			b.WriteString(c.Name)
		}
		b.WriteString(". All Rights Reserved.")
		return b.String()
	}
	return ""
}

// provides a nice string representation of the Copyright statement
func (c *CopyrightResult) String() string {
	return c.Stringify("©")
}

// provides a nice string representation of the Copyright statement
func (c *CopyrightResult) HTMLString() string {
	return c.Stringify("&copy;")
}

var re = regexp.MustCompile("(?i)(copyright)\\s*(\\(c\\)|&copy;|&#169;|&#xa9;|©)?\\s*(\\d{4})\\s*[-,]?\\s*(\\d{4})?\\s*(by)?\\s*(.*)")

func cleanupName(name string) string {
	name = strings.Replace(name, "All rights reserved", "", -1)
	name = strings.Replace(name, " .", "", -1)
	name = strings.TrimSpace(name)
	if strings.HasSuffix(name, ".") {
		name = name[0 : len(name)-1]
	}
	return name
}

// ParseCopyright will attempt to parse copyright details from a line (if found)
func ParseCopyright(line string) (*CopyrightResult, error) {
	result := &CopyrightResult{}
	if re.MatchString(line) {
		tokens := re.FindStringSubmatch(line)
		var date1, date2, name string
		len := len(tokens)
		if len > 3 {
			date1 = tokens[3]
		}
		if len > 4 {
			date2 = tokens[4]
		}
		if len > 5 {
			name = tokens[6]
		}
		result.Found = true
		if date1 != "" {
			result.Dates = make([]string, 0)
			result.Dates = append(result.Dates, date1)
			if date2 != "" {
				result.Dates = append(result.Dates, date2)
			}
		}
		result.Name = cleanupName(name)
	}
	return result, nil
}
