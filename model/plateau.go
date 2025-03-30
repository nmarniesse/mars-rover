package model

import (
	"errors"
)

type Plateau struct {
	maxX   int
	maxY   int
	rovers []*Rover
}

func CreatePlateau(maxX int, maxY int) (*Plateau, error) {
	if maxX < 0 {
		return nil, errors.New("maxX cannot be negative")
	}
	if maxY < 0 {
		return nil, errors.New("maxY cannot be negative")
	}

	return &Plateau{maxX, maxY, make([]*Rover, 0)}, nil
}

func (plateau *Plateau) MaxX() int {
	return plateau.maxX
}

func (plateau *Plateau) MaxY() int {
	return plateau.maxY
}

func (plateau *Plateau) Rovers() []*Rover {
	return plateau.rovers
}

func (plateau *Plateau) AddRover(x int, y int, direction string) (*Rover, error) {
	if x > plateau.maxX {
		return nil, errors.New("rover has not landed on the plateau")
	}
	if y > plateau.maxY {
		return nil, errors.New("rover has not landed on the plateau")
	}

	rover, err := CreateRover(x, y, direction, plateau)
	if err != nil {
		return nil, err
	}

	plateau.rovers = append(plateau.rovers, rover)

	return rover, nil
}
