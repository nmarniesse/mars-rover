package command

import (
	"fmt"

	"github.com/nmarniesse/mars-rover/model"
)

func ApplyInstructions(plateau *model.Plateau, rover *model.Rover, instructions string) error {
	for _, instruction := range instructions {
		switch instruction {
		case 'M':
			rover.MoveRoverForward(plateau)
		case 'L':
			rover.RotateLeft()
		case 'R':
			rover.RotateRight()
		default:
			return fmt.Errorf("unknown [%c] instruction", instruction)
		}
	}

	return nil
}
