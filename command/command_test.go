package command

import (
	"testing"

	"github.com/nmarniesse/mars-rover/model"
	"github.com/stretchr/testify/assert"
)

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
