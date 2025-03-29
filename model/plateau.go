package model

import (
	"errors"
)

type Plateau struct {
	MaxX int
	MaxY int
}

func CreatePlateau(maxX int, maxY int) (*Plateau, error) {
	if maxX <= 0 {
		return nil, errors.New("maxX cannot be equal or lower than 0")
	}
	if maxY <= 0 {
		return nil, errors.New("maxY cannot be equal or lower than 0")
	}

	return &Plateau{maxX, maxY}, nil
}
