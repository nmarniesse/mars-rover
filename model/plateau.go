package model

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Plateau struct {
	MaxX   int `validate:"gte=0"`
	MaxY   int `validate:"gte=0"`
	Rovers []*Rover
}

func CreatePlateau(maxX int, maxY int) (*Plateau, error) {
	var rovers []*Rover
	plateau := Plateau{maxX, maxY, rovers}
	validate := validator.New()
	err := validate.Struct(plateau)

	return &plateau, err
}

func (plateau *Plateau) AddRover(x int, y int, direction string) (*Rover, error) {
	if x > plateau.MaxX {
		return nil, errors.New("Rover has not landed on the plateau")
	}
	if y > plateau.MaxY {
		return nil, errors.New("Rover has not landed on the plateau")
	}

	rover, err := CreateRover(x, y, direction, plateau)
	if err != nil {
		return nil, err
	}

	plateau.Rovers = append(plateau.Rovers, rover)

	return rover, nil
}
