package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveAccents(t *testing.T) {
	stringWithAccents := "áéíóúàèìòùâêîôûãõçñ"
	stringNormalized := "aeiouaeiouaeiouaocn"

	stringNormalizedResult, err := RemoveAccents(stringWithAccents)

	assert.Nil(t, err)
	assert.Equal(t, stringNormalized, stringNormalizedResult)
}
