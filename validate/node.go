//
// node.go
// Copyright (C) 2016 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package validate

import (
	"fmt"
	"reflect"
	"regexp"
)

var (
	yamlKey  = regexp.MustCompile(`^ *-? ?(?P<key>.*):`)
	yamlElem = regexp.MustCompile(`^ *-`)
)

type node struct {
	name     string
	line     int
	children []node
	field    reflect.StructField
	reflect.Value
}

// Child attempts to find the child with the given name in the node's list of
// children. If no such child is found, an invalid node is returned.
func (n node) Child(name string) node {
	for _, c := range n.children {
		if c.name == name {
			return c
		}
	}

	return node{}
}

// HummanType returns the human-consumable string representation of the type of
// the node
func (n node) HumanType() string {
	switch k := n.Kind(); k {
	case reflect.Slice:
		c := n.Type().Elem()
		return "[]" + node{Value: reflect.New(c).Elem()}.HumanType()
	default:
		return k.String()
	}
}

// NewNode returns the node representation of the given value. The context
// will be used in an attempt to determine line numbers for the given value
func NewNode(value interface{}, context context) node {
	var n node
	toNode(value, context, &n)
	return n
}

// toNode converts the given value into a node and then recurisively processes
// each of the nodes components (e.g. fields, array elements, keys)
func toNode(v interface{}, c context, n *node) {
	vv := reflect.ValueOf(v)
	if !vv.IsValid() {
		return
	}

	n.Value = vv
	switch vv.Kind() {
	case reflect.Struct:
		// Walk over each filed in the structure, skipping unexported fields,
		// and create a node for it
		for i := 0; i < vv.Type().NumField(); i++ {
			ft := vv.Type().Field(i)
			k := ft.Tag.Get("yaml")
			if k == "-" || k == "" {
				continue
			}

			cn := node{name: k, field: ft}
			c, ok := findKey(cn.name, c)
			if ok {
				cn.line = c.lineNumber
			}
			toNode(vv.Field(i).Interface(), c, &cn)
			n.children = append(n.children, cn)
		}
	case reflect.Map:
		// Walk over each key in the map and create a node for it
		v := v.(map[interface{}]interface{})
		for k, cv := range v {
			cn := node{name: fmt.Sprintf("%s", k)}
			c, ok := findKey(cn.name, c)
			if ok {
				cn.line = c.lineNumber
			}
			toNode(cv, c, &cn)
			n.children = append(n.children, cn)
		}
	case reflect.Slice:
		// Walk over each element in the slice and create a node for it
		// While iterating over the slice, preserve the context after it
		// is modified. This allows the line numbers to reflect the content
		// element instead of the first
		for i := 0; i < vv.Len(); i++ {
			cn := node{
				name:  fmt.Sprintf("%s[%d]", n.name, i),
				field: n.field,
			}
			var ok bool
			c, ok := findElem(c)
			if ok {
				cn.line = c.lineNumber
			}
			toNode(vv.Index(i).Interface(), c, &cn)
			n.children = append(n.children, cn)
			c.Increment()
		}
	case reflect.String, reflect.Int, reflect.Bool, reflect.Float64:
	default:
		panic(fmt.Sprintf("toNode(): unhandled kind %s", vv.Kind()))
	}
}

// findKey attempts to find the requested key within the provided context.
// A modified copy of the context is returned with every line up to the key
// incremented past. A boolean, true if the key was found, is also returned.
func findKey(key string, context context) (context, bool) {
	return find(yamlKey, key, context)
}

// findElem attempts to find an array element within the provided context.
// A modified copy of the context is returned with every line up to the array
// element incremented past. A boolean, true if the Elem was found, is also
// returned.
func findElem(context context) (context, bool) {
	return find(yamlElem, "", context)
}

func find(exp *regexp.Regexp, key string, context context) (context, bool) {
	for len(context.currentLine) > 0 || len(context.remainingLines) > 0 {
		matches := exp.FindStringSubmatch(context.currentLine)
		if len(matches) > 0 && (key == "" || matches[1] == key) {
			return context, true
		}
		context.Increment()
	}

	return context, false
}
