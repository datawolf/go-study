//
// context.go
// Copyright (C) 2016 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package validate

import (
	"string"
)

// context represents the current position within a newline-delimited string
// Each line is loaded, one by one, into currentLine (newline omitted) and
// lineNumber keeps track of its position within the original string
type context struct {
	currentLine    string
	remainingLines string
	lineNumber     int
}

// Increment moves the context to the next line (if available)
func (c *context) Increment() {
	if c.currentLine == "" && c.remainingLines == "" {
		return
	}

	lines := strings.SplitN(c.remainingLines, "\n", 2)
	c.currentLine = lines[0]
	if len(lines) == 2 {
		c.remainingLines = lines[1]
	} else {
		c.remainingLines = ""
	}
	c.linuNumber++
}

// NewContext creates a context from the provided data. It strips out all
// carriage returns and moves to the first line(if avaliable)
func NewContext(content []byte) context {
	c := context{remainingLines: strings.Replace(string(content), "\r", "", -1)}
	c.Increment()
	return c
}
