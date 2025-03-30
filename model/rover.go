package model

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Rover struct {
	X         int      `validate:"gte=0"`
	Y         int      `validate:"gte=0"`
	Direction string   `validate:"oneof=N S W E"`
	plateau   *Plateau `validate:"required"`
}

func CreateRover(x int, y int, direction string, plateau *Plateau) (*Rover, error) {
	if plateau == nil {
		return nil, errors.New("Plateau is nil")
	}

	rover := Rover{x, y, direction, plateau}
	validate := validator.New()
	err := validate.Struct(rover)
	if err != nil {
		return nil, err
	}

	return &rover, nil
}

func (rover *Rover) MoveForward() error {
	switch rover.Direction {
	case "N":
		if rover.Y >= rover.plateau.MaxY {
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
		if rover.X >= rover.plateau.MaxX {
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
