//
// context_test.go
// Copyright (C) 2016 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package validate

import (
	"reflect"
	"testing"
)

func TestNewContext(t *testing.T) {
	tests := []struct {
		in  string
		out context
	}{
		{
			out: context{
				currentLine:    "",
				remainingLines: "",
				lineNumber:     0,
			},
		},
		{
			in: "this\r\nis\r\na\r\ntest",
			out: context{
				currentLine:    "this",
				remainingLines: "is\na\ntest",
				lineNumber:     1,
			},
		},
	}

	for _, tt := range tests {
		if out := NewContext([]byte(tt.in)); !reflect.DeepEqual(tt.out, out) {
			t.Errorf("bad context (%q): want %#v, got %#v", tt.in, tt.out, out)
		}
	}
}

func TestIncrement(t *testing.T) {
	tests := []struct {
		init context
		op   func(c *context)
		res  context
	}{
		{
			init: context{
				currentLine:    "",
				remainingLines: "",
				lineNumber:     0,
			},
			res: context{
				currentLine:    "",
				remainingLines: "",
				lineNumber:     0,
			},
			op: func(c *context) {
				c.Increment()
			},
		},
		{
			init: context{
				currentLine:    "test",
				remainingLines: "",
				lineNumber:     1,
			},
			res: context{
				currentLine:    "",
				remainingLines: "",
				lineNumber:     2,
			},
			op: func(c *context) {
				c.Increment()
				c.Increment()
				c.Increment()
			},
		},
		{
			init: context{
				currentLine:    "this",
				remainingLines: "is\na\ntest",
				lineNumber:     1,
			},
			res: context{
				currentLine:    "is",
				remainingLines: "a\ntest",
				lineNumber:     2,
			},
			op: func(c *context) {
				c.Increment()
			},
		},
		{
			init: context{
				currentLine:    "this",
				remainingLines: "is\na\ntest",
				lineNumber:     1,
			},
			res: context{
				currentLine:    "test",
				remainingLines: "",
				lineNumber:     4,
			},
			op: func(c *context) {
				c.Increment()
				c.Increment()
				c.Increment()
			},
		},
	}

	for i, tt := range tests {
		res := tt.init
		if tt.op(&res); !reflect.DeepEqual(tt.res, res) {
			t.Errorf("bad context (%d, %#v): want %#v, got %#v", i, tt.init, tt.res, res)
		}
	}

}
