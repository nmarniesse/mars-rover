package model

import (
	"errors"
	"strconv"
	"strings"
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

func CreateRoverFromLine(line string) (*Rover, error) {
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

	return CreateRover(x, y, elements[2])
}

func (rover *Rover) MoveForward() {
	if rover.Direction == "N" {
		rover.Y++
	} else if rover.Direction == "S" {
		rover.Y--
	} else if rover.Direction == "W" {
		rover.X--
	} else if rover.Direction == "E" {
		rover.X++
	}
}
