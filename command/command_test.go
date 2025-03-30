package command

import (
	"testing"

	"github.com/nmarniesse/mars-rover/model"
	"github.com/stretchr/testify/assert"
)

func TestItCreatesAPlateauFromLine(t *testing.T) {
	p, err := CreatePlateauFromLine("4 5")
	assert.Nil(t, err)

	assert.Equal(t, 4, p.MaxX())
	assert.Equal(t, 5, p.MaxY())
}

func TestItAddsARoverFromLine(t *testing.T) {
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

	plateau, _ := model.CreatePlateau(5, 5)
	for _, c := range cases {
		rover, err := AddRoverFromLine(plateau, c.line)
		assert.Nil(t, err)

		assert.Equal(t, c.expectedX, rover.X())
		assert.Equal(t, c.expectedY, rover.Y())
		assert.Equal(t, c.expectedDirection, rover.Direction())
	}
}

func TestItCannotAddARoverFromLine(t *testing.T) {
	plateau, _ := model.CreatePlateau(4, 5)
	rover, err := AddRoverFromLine(plateau, "5 2 N")
	assert.Nil(t, rover)
	assert.Error(t, err)
}

func TestItAppliesInstructionsToARover(t *testing.T) {
	plateau, err := model.CreatePlateau(5, 5)
	assert.Nil(t, err)

	rover1, _ := model.CreateRover(1, 2, "N", plateau)
	rover2, _ := model.CreateRover(3, 3, "E", plateau)

	cases := []struct {
		rover              *model.Rover
		instructions       string
		expectedX          int
		expectedY          int
		expectedDirections string
	}{
		{rover1, "LMLMLMLMM", 1, 3, "N"},
		{rover2, "MMRMMRMRRM", 5, 1, "E"},
	}

	for _, c := range cases {
		err = ApplyInstructions(c.rover, c.instructions)
		assert.Nil(t, err)

		assert.Equal(t, c.expectedX, c.rover.X())
		assert.Equal(t, c.expectedY, c.rover.Y())
		assert.Equal(t, c.expectedDirections, c.rover.Direction())
	}
}
