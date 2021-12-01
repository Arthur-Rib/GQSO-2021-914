package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_soma(t *testing.T) {
	result, err := soma("1", "2")
	assert.NoError(t, err)
	assert.Equal(t, result, 3.0)
}
