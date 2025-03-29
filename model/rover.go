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

func (rover *Rover) MoveRoverForward(plateau *Plateau) error {
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
