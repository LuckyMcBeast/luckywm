package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandleArgVersion(t *testing.T) {
	fakeArgs := []string{"application", "-v"}

	assert.Equal(t, getVersion(), handleArgs(fakeArgs))
}

func TestHandleInvalidArg(t *testing.T){
	fakeArgs := []string{ "application", "asdfasdfasdfa"}

	assert.Equal(t, getUsage(), handleArgs(fakeArgs))
}