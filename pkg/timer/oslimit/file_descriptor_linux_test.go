package oslimits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileDescriptors(t *testing.T) {
	fd, err := GetOSSimultaneousFileDescriptors()

	assert.NoError(t, err)
	assert.Greater(t, fd, uint64(0))
}

func TestSetMaxFileDescriptors(t *testing.T) {
	fdInit, _ := GetOSSimultaneousFileDescriptors()

	err := SetMaxOSSimultaneousFileDescriptors()

	fdFinal, _ := GetOSSimultaneousFileDescriptors()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, fdFinal, fdInit)
}
