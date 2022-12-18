package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalesog/go-pix-utils/emv"
)

func TestParsePixStatic(t *testing.T) {
	testingEmv, err := emv.ParsePix(staticEmv)

	assert.Nil(t, err)
	assert.Equal(t, staticElements, testingEmv.Elements)
	assert.Equal(t, staticEmv, testingEmv.Raw)
	assert.Equal(t, "static", testingEmv.Type)
}

func TestParsePixDynamic(t *testing.T) {
	testingEmv, err := emv.ParsePix(dynamicEmv)

	assert.Nil(t, err)
	assert.Equal(t, dynamicElements, testingEmv.Elements)
	assert.Equal(t, dynamicEmv, testingEmv.Raw)
	assert.Equal(t, "dynamic", testingEmv.Type)
}

func TestDetectInvalidCrc(t *testing.T) {
	_, err := emv.ParsePix(staticEmvWrongCrc)
	assert.Errorf(t, err, "invalid crc")
}
