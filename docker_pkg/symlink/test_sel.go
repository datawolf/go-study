//
// test_sel.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type linkOrDir struct {
	path   string
	target string
}

func makeFs(tmpdir string, fs []linkOrDir) error {
	for _, s := range fs {
		s.path = filepath.Join(tmpdir, s.path)
		if s.target == "" {
			os.MkdirAll(s.path, 0755)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(s.path), 0755); err != nil {
			return err
		}
		if err := os.Symlink(s.target, s.path); err != nil && !os.IsExist(err) {
			return err
		}
	}

	return nil
}

func main() {
	tmpdir, err := ioutil.TempDir("", "Main")
	if err != nil {
		fmt.Println(err)
	}
	defer os.RemoveAll(tmpdir)

	if err := makeFs(tmpdir, []linkOrDir{
		{path: "linkdir", target: "realdir"},
		{path: "linkdir/foo/bar"},
	}); err != nil {
		fmt.Println(err)
	}
}
