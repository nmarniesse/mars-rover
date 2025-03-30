package model

import (
	"errors"
)

type Rover struct {
	x         int
	y         int
	direction string
	plateau   *Plateau
}

func CreateRover(x int, y int, direction string, plateau *Plateau) (*Rover, error) {
	if x < 0 {
		return nil, errors.New("x cannot be negative")
	}

	if y < 0 {
		return nil, errors.New("y cannot be negative")
	}

	if direction != "N" && direction != "S" && direction != "W" && direction != "E" {
		return nil, errors.New("direction is not valid")
	}

	if plateau == nil {
		return nil, errors.New("Plateau is not defined")
	}

	return &Rover{x, y, direction, plateau}, nil
}

func (rover *Rover) X() int {
	return rover.x
}

func (rover *Rover) Y() int {
	return rover.y
}

func (rover *Rover) Direction() string {
	return rover.direction
}

func (rover *Rover) MoveForward() error {
	switch rover.direction {
	case "N":
		if rover.y >= rover.plateau.maxY {
			return errors.New("out of bound: cannot move the rover forward")
		}
		rover.y++
	case "S":
		if rover.y <= 0 {
			return errors.New("out of bound: cannot move the rover forward")
		}
		rover.y--
	case "W":
		if rover.x <= 0 {
			return errors.New("out of bound: cannot move the rover forward")
		}
		rover.x--
	case "E":
		if rover.x >= rover.plateau.maxX {
			return errors.New("out of bound: cannot move the rover forward")
		}
		rover.x++
	}

	return nil
}

func (rover *Rover) RotateLeft() {
	switch rover.direction {
	case "N":
		rover.direction = "W"
	case "S":
		rover.direction = "E"
	case "W":
		rover.direction = "S"
	case "E":
		rover.direction = "N"
	}
}

func (rover *Rover) RotateRight() {
	switch rover.direction {
	case "N":
		rover.direction = "E"
	case "S":
		rover.direction = "W"
	case "W":
		rover.direction = "N"
	case "E":
		rover.direction = "S"
	}
}
