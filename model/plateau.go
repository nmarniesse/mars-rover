package model

import "github.com/go-playground/validator/v10"

type Plateau struct {
	MaxX int `validate:"gte=0"`
	MaxY int `validate:"gte=0"`
}

func CreatePlateau(maxX int, maxY int) (*Plateau, error) {
	plateau := Plateau{maxX, maxY}
	validate := validator.New()
	err := validate.Struct(plateau)

	return &plateau, err
}
