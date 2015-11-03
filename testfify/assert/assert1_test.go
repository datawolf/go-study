//
// assert_test.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	var a string = "Hello"
	var	b string = "Hello"
	var object string = "Something"

	assert := assert.New(t)

	assert.Equal(a, b, "The two words should be the same.")

	// assert equality
	assert.Equal(123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(object)

	// assert for not nil (good when you expect something)
	if assert.NotNil(object) {
		// now we know that object is not nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("Something", object)
	}
}
