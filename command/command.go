package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/nmarniesse/mars-rover/model"
)

func CreatePlateauFromLine(line string) (*model.Plateau, error) {
	maxCoordinates := strings.Split(line, " ")
	if len(maxCoordinates) != 2 {
		return nil, errors.New("error while creating plateau from line: expecting exactly 2 elements")
	}

	maxX, err := strconv.Atoi(maxCoordinates[0])
	if err != nil {
		return nil, errors.New("error while creating plateau from line: first element not an integer")
	}

	maxY, err := strconv.Atoi(maxCoordinates[1])
	if err != nil {
		return nil, errors.New("error while creating plateau from line: second element not an integer")
	}

	return model.CreatePlateau(maxX, maxY)
}

func AddRoverFromLine(plateau *model.Plateau, line string) (*model.Rover, error) {
	elements := strings.Split(line, " ")
	if len(elements) != 3 {
		return nil, errors.New("error while creating rover from line: expecting exactly 3 elements")
	}

	x, err := strconv.Atoi(elements[0])
	if err != nil {
		return nil, errors.New("error while creating rover from line: first element not an integer")
	}

	y, err := strconv.Atoi(elements[1])
	if err != nil {
		return nil, errors.New("error while creating rover from line: second element not an integer")
	}

	return plateau.AddRover(x, y, elements[2])
}

func ApplyInstructions(rover *model.Rover, instructions string) error {
	for _, instruction := range instructions {
		switch instruction {
		case 'M':
			rover.MoveForward()
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
