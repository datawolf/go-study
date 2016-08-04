//
// rules.go
// Copyright (C) 2016 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package validate

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type rule func(config node, report *Report)

// Rules contains all of the validation rules
var Rules []rule = []rule{
	checkStructure,
	checkValidity,
}

// checkStructure compares the provided config to the empty config.CloudConfig
// structure. Each node is checked to make sure that it exists in the known
// structure and that its type is compatible.
func checkStructure(cfg node, report *Report) {
	g := NewNode(config.CloudConfig{}, NewContext([]byte{}))
	checkNodeStructure(cfg, g, report)
}

func checkNodeStructure(n, g node, r *Report) {
	if !isCompatible(n.Kind(), g.Kind()) {
		r.Warning(n.line, fmt.Sprintf("incorrect type for %q (want %s)", n.name, g.HumanType()))
		return
	}

	switch g.Kind() {
	case reflect.Struct:
		for _, cn := range n.children {
			if cg := g.Child(cn.name); cg.IsValid() {
				if msg := cg.field.Tag.Get("deprecated"); msg != "" {
					r.Warning(cn.line, fmt.Sprintf("deprecated key %q (%s)", cn.name, msg))
				}
				checkNodeStructure(cn, cg, r)
			} else {
				r.Warning(cn.line, fmt.Sprintf("unrecognized key %q", cn.name))
			}
		}
	case reflect.Slice:
		for _, cn := range n.children {
			var cg node
			c := g.Type().Elem()
			toNode(reflect.New(c).Elem().Interface(), context{}, &cg)
			checkNodeStructure(cn, cg, r)
		}
	case reflect.String, reflect.Int, reflect.Float64, reflect.Bool:
	default:
		panic(fmt.Sprintf("checkNodeStructure(): unhandled kind %s", g.Kind()))
	}
}

// isCompatible determines if the type of kind n can be converted to the type of
// kind g in the content of YAML. This is not an exhaustive list, but its enough for
// the purposes of cloud-config validation.
func isCompatible(n, g reflect.Kind) bool {
	switch g {
	case reflect.String:
		return n == reflect.String || n == reflect.Int || n == reflect.Float64 || n == reflect.Bool
	case reflect.Struct:
		return n == reflect.Struct || n == reflect.Map
	case reflect.Float64:
		return n == reflect.Float64 || n == reflect.Int
	case reflect.Bool, reflect.Slice, reflect.Int:
		return n == g
	default:
		panic(fmt.Sprintf("isCompatible(): unhandled kind %s", g))
	}
}

// checkValidity check the value of every node in the provided config by AssertValid on it.
func checkValidity(cfg node, report *Report) {
	g := NewNode(config.CloudConfig{}, NewContext([]byte{}))
	checkNodeValidity(cfg, g, report)
}

func checkNodeValidity(n, g node, r *Report) {
	if err := assertValid(n.Value, g.field.Tag.Get("valid")); err != nil {
		r.Error(n.line, fmt.Sprintf("invalid value %v", n.Value.Interface()))
	}

	switch g.Kind() {
	case reflect.Struct:
		for _, cn := range n.children {
			if cg := g.Child(cn.name); cg.IsValid() {
				checkNodeValidity(cn, cg, r)
			}
		}
	case reflect.Slice:
		for _, cn := range n.children {
			var cg node
			c := g.Type().Elem()
			toNode(reflect.New(c).Elem().Interface(), context{}, &cg)
			checkNodeValidity(cn, cg, r)
		}
	case reflect.String, reflect.Int, reflect.Float64, reflect.Bool:
	default:
		panic(fmt.Sprintf("checkNodeValidity(): unhandled kind %s", g.Kind()))
	}
}

func assertValid(value reflect.Value, valid string) error {
	if valid == "" || isZero(value) {
		return nil
	}

	vs := fmt.Sprintf("%v", value.Interface())
	if m, _ := regexp.MatchString(valid, vs); m {
		return nil
	}

	return errors.New(fmt.Sprintf("invalid value %q for option (valid options: %q)", vs, valid))
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Struct:
		vt := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if isFieldExported(vt.Field(i)) && !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	}
}

func isFieldExported(f reflect.StructField) bool {
	return f.PkgPath == ""
}
