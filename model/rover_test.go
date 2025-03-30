package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCreatesARover(t *testing.T) {
	plateau, _ := CreatePlateau(5, 5)
	rover, err := CreateRover(1, 2, "N", plateau)
	assert.Nil(t, err)

	assert.Equal(t, 1, rover.x)
	assert.Equal(t, 2, rover.y)
	assert.Equal(t, "N", rover.direction)
}

func TestItCannotCreateARoverWithWrongDirection(t *testing.T) {
	plateau, _ := CreatePlateau(5, 5)
	_, err := CreateRover(1, 2, "unknown", plateau)
	assert.Error(t, err)
}

func TestItCannotCreateARoverWithNegativeX(t *testing.T) {
	plateau, _ := CreatePlateau(5, 5)
	rover, err := CreateRover(-1, 2, "N", plateau)
	assert.Nil(t, rover)
	assert.Error(t, err)
}

func TestItCannotCreateARoverWithNegativeY(t *testing.T) {
	plateau, _ := CreatePlateau(5, 5)
	_, err := CreateRover(1, -2, "N", plateau)
	assert.Error(t, err)
}

func TestItCannotCreateARoverWithoutPlateau(t *testing.T) {
	_, err := CreateRover(1, 2, "N", nil)
	assert.Error(t, err)
}

func TestItMovesARoverForward(t *testing.T) {
	plateau, _ := CreatePlateau(4, 5)
	cases := []struct {
		rover     *Rover
		expectedX int
		expectedY int
	}{
		{&Rover{2, 2, "N", plateau}, 2, 3},
		{&Rover{2, 2, "S", plateau}, 2, 1},
		{&Rover{2, 2, "W", plateau}, 1, 2},
		{&Rover{2, 2, "E", plateau}, 3, 2},
	}

	for _, c := range cases {
		err := c.rover.MoveForward()
		assert.Nil(t, err)

		assert.Equal(t, c.expectedX, c.rover.x)
		assert.Equal(t, c.expectedY, c.rover.y)
	}
}
