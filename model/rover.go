package model

import (
	"errors"
)

type Rover struct {
	X         int
	Y         int
	Direction string
}

func CreateRover(x int, y int, direction string) (*Rover, error) {
	if x < 0 {
		return nil, errors.New("x cannot be lower than 0")
	}

	if y < 0 {
		return nil, errors.New("y cannot be lower than 0")
	}

	if direction != "N" && direction != "S" && direction != "W" && direction != "E" {
		return nil, errors.New("direction should be N, S, W or E")
	}

	return &Rover{x, y, direction}, nil
}

func (rover *Rover) MoveForward(plateau *Plateau) error {
	switch rover.Direction {
	case "N":
		if rover.Y >= plateau.MaxY {
			return errors.New("out of bound: cannot move the rover forward")
		}
		rover.Y++
	case "S":
		if rover.Y <= 0 {
			return errors.New("out of bound: cannot move the rover forward")
		}
		rover.Y--
	case "W":
		if rover.X <= 0 {
			return errors.New("out of bound: cannot move the rover forward")
		}
		rover.X--
	case "E":
		if rover.X >= plateau.MaxX {
			return errors.New("out of bound: cannot move the rover forward")
		}
		rover.X++
	}

	return nil
}

func (rover *Rover) RotateLeft() {
	switch rover.Direction {
	case "N":
		rover.Direction = "W"
	case "S":
		rover.Direction = "E"
	case "W":
		rover.Direction = "S"
	case "E":
		rover.Direction = "N"
	}
}

func (rover *Rover) RotateRight() {
	switch rover.Direction {
	case "N":
		rover.Direction = "E"
	case "S":
		rover.Direction = "W"
	case "W":
		rover.Direction = "N"
	case "E":
		rover.Direction = "S"
	}
}
