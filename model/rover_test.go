package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCreatesARover(t *testing.T) {
	rover, err := CreateRover(1, 2, "N")
	assert.Nil(t, err)

	assert.Equal(t, 1, rover.X)
	assert.Equal(t, 2, rover.Y)
	assert.Equal(t, "N", rover.Direction)
}

func TestItCreatesARoverFromLine(t *testing.T) {
	cases := []struct {
		line              string
		expectedX         int
		expectedY         int
		expectedDirection string
	}{
		{"2 2 N", 2, 2, "N"},
		{"1 2 S", 1, 2, "S"},
		{"2 5 W", 2, 5, "W"},
		{"0 2 E", 0, 2, "E"},
	}

	for _, c := range cases {
		rover, err := CreateRoverFromLine(c.line)
		assert.Nil(t, err)

		assert.Equal(t, c.expectedX, rover.X)
		assert.Equal(t, c.expectedY, rover.Y)
		assert.Equal(t, c.expectedDirection, rover.Direction)
	}
}

func TestItCannotCreateARoverWithWrongDirection(t *testing.T) {
	_, err := CreateRover(1, 2, "unknown")
	assert.Error(t, err)
}

func TestItCannotCreateARoverWithNegativeX(t *testing.T) {
	_, err := CreateRover(-1, 2, "N")
	assert.Error(t, err)
}

func TestItCannotCreateARoverWithNegativeY(t *testing.T) {
	_, err := CreateRover(1, -2, "N")
	assert.Error(t, err)
}

func TestItMovesARoverForward(t *testing.T) {
	cases := []struct {
		rover     *Rover
		plateau   *Plateau
		expectedX int
		expectedY int
	}{
		{&Rover{2, 2, "N"}, &Plateau{4, 5}, 2, 3},
		{&Rover{2, 2, "S"}, &Plateau{4, 5}, 2, 1},
		{&Rover{2, 2, "W"}, &Plateau{4, 5}, 1, 2},
		{&Rover{2, 2, "E"}, &Plateau{4, 5}, 3, 2},
	}

	for _, c := range cases {
		err := c.rover.MoveRoverForward(c.plateau)
		assert.Nil(t, err)

		assert.Equal(t, c.expectedX, c.rover.X)
		assert.Equal(t, c.expectedY, c.rover.Y)
	}
}
