//
// report_test.go
// Copyright (C) 2016 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package validate

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEntry(t *testing.T) {
	tests := []struct {
		entry Entry
		str   string
		json  []byte
	}{
		{
			Entry{entryInfo, "test info", 1},
			"line 1: info: test info",
			[]byte(`{"kind":"info","line":1,"message":"test info"}`),
		},
		{
			Entry{entryWarning, "test warning", 2},
			"line 2: warning: test warning",
			[]byte(`{"kind":"warning","line":2,"message":"test warning"}`),
		},
		{
			Entry{entryError, "test error", 12},
			"line 12: error: test error",
			[]byte(`{"kind":"error","line":12,"message":"test error"}`),
		},
	}

	for _, tt := range tests {
		if str := tt.entry.String(); tt.str != str {
			t.Errorf("bad string(%q): want %q, got %q", tt.entry, tt.str, str)
		}
		json, err := tt.entry.MarshalJSON()
		if err != nil {
			t.Errorf("bad error(%q): want %q, got %q", tt.entry, nil, err)
		}
		if !bytes.Equal(tt.json, json) {
			t.Errorf("bad JSON(%q): want %q, got %q", tt.entry, tt.json, json)
		}
	}
}

func TestReport(t *testing.T) {
	type reportFunc struct {
		fn      func(*Report, int, string)
		line    int
		message string
	}

	tests := []struct {
		fs []reportFunc
		es []Entry
	}{
		{
			[]reportFunc{
				{(*Report).Warning, 1, "test warning 1"},
				{(*Report).Error, 2, "test error 2"},
				{(*Report).Info, 10, "test info 10"},
			},
			[]Entry{
				{entryWarning, "test warning 1", 1},
				{entryError, "test error 2", 2},
				{entryInfo, "test info 10", 10},
			},
		},
	}

	for _, tt := range tests {
		r := Report{}
		for _, f := range tt.fs {
			f.fn(&r, f.line, f.message)
		}

		if es := r.Entries(); !reflect.DeepEqual(tt.es, es) {
			t.Errorf("bad entries (%v):\nwant %#v, \ngot %#v", tt.fs, tt.es, es)
		}
	}
}
