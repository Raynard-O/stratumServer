package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomHex(t *testing.T) {
	s := randomHex(19)
	assert.IsType(t, string(""), s)
}

