package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCreatesAPlateau(t *testing.T) {
	p, err := CreatePlateau(4, 5)
	assert.Nil(t, err)

	assert.Equal(t, 4, p.maxX)
	assert.Equal(t, 5, p.maxY)
	assert.Equal(t, 0, len(p.rovers))
}

func TestItCannotCreateAPlateauWithNegativeX(t *testing.T) {
	_, err := CreatePlateau(-4, 5)
	assert.Error(t, err)
}

func TestItCannotCreateAPlateauWithNegativeY(t *testing.T) {
	_, err := CreatePlateau(4, -2)
	assert.Error(t, err)
}

func TestItAddsARoverOnPlateau(t *testing.T) {
	plateau, _ := CreatePlateau(4, 5)
	rover, err := plateau.AddRover(4, 5, "S")

	assert.Nil(t, err)
	assert.Equal(t, 4, rover.x)
	assert.Equal(t, 1, len(plateau.rovers))
}

func TestItCannotAddARoverOutsideThePlateauOnXAxis(t *testing.T) {
	plateau, _ := CreatePlateau(4, 5)
	rover, err := plateau.AddRover(5, 5, "S")

	assert.Error(t, err)
	assert.Nil(t, rover)
	assert.Equal(t, 0, len(plateau.rovers))
}


func TestItCannotAddARoverOutsideThePlateauOnYAxis(t *testing.T) {
	plateau, _ := CreatePlateau(4, 5)
	rover, err := plateau.AddRover(1, 6, "S")

	assert.Error(t, err)
	assert.Nil(t, rover)
	assert.Equal(t, 0, len(plateau.rovers))
}
