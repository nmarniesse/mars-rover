package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCreatesAPlateau(t *testing.T) {
	p, err := CreatePlateau(4, 5)
	assert.Nil(t, err)

	assert.Equal(t, 4, p.MaxX)
	assert.Equal(t, 5, p.MaxY)
}

func TestItCannotCreateAPlateauWithNegativeX(t *testing.T) {
	_, err := CreatePlateau(-4, 5)
	assert.Error(t, err)
}
