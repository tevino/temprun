package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDstNotClosed(t *testing.T) {
	tempfile, err := ioutil.TempFile("", "temprun_test")
	assert.NoError(t, err)
	defer os.Remove(tempfile.Name())
	t.Logf("Using tempfile: %s", tempfile.Name())

	dst, err := openDst(tempfile.Name(), 0755)
	assert.NoError(t, err)
	assert.NoError(t, dst.Close())
}
