package model

import (
	"errors"
	"strconv"
	"strings"
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

func CreatePlateauFromLine(line string) (*Plateau, error) {
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

	return CreatePlateau(maxX, maxY)
}
