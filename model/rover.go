package model

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Rover struct {
	X         int    `validate:"gte=0"`
	Y         int    `validate:"gte=0"`
	Direction string `validate:"oneof=N S W E"`
}

func CreateRover(x int, y int, direction string) (*Rover, error) {
	rover := Rover{x, y, direction}
	validate := validator.New()
	err := validate.Struct(rover)

	return &rover, err
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
