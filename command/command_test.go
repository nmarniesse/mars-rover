package command

import (
	"testing"

	"github.com/nmarniesse/mars-rover/model"
	"github.com/stretchr/testify/assert"
)

func TestItCreatesAPlateauFromLine(t *testing.T) {
	p, err := CreatePlateauFromLine("4 5")
	assert.Nil(t, err)

	assert.Equal(t, 4, p.MaxX)
	assert.Equal(t, 5, p.MaxY)
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

func TestItAppliesInstructionsToARover(t *testing.T) {
	plateau, err := model.CreatePlateau(5, 5)
	assert.Nil(t, err)

	cases := []struct {
		plateau            *model.Plateau
		rover              *model.Rover
		instructions       string
		expectedX          int
		expectedY          int
		expectedDirections string
	}{
		{plateau, &model.Rover{X: 1, Y: 2, Direction: "N"}, "LMLMLMLMM", 1, 3, "N"},
		{plateau, &model.Rover{X: 3, Y: 3, Direction: "E"}, "MMRMMRMRRM", 5, 1, "E"},
	}

	for _, c := range cases {
		err = ApplyInstructions(c.plateau, c.rover, c.instructions)
		assert.Nil(t, err)

		assert.Equal(t, c.expectedX, c.rover.X)
		assert.Equal(t, c.expectedY, c.rover.Y)
		assert.Equal(t, c.expectedDirections, c.rover.Direction)
	}
}
