package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nmarniesse/mars-rover/command"
	"github.com/nmarniesse/mars-rover/model"
)

func main() {
	filename := "./input/challenge.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	lines := getLinesFromFile(filename)
	firstLine := lines[0]
	plateau, err := model.CreatePlateauFromLine(firstLine)
	if err != nil {
		log.Fatal(err)
	}

	var rovers [](*model.Rover)
	for i := 1; i < len(lines); i = i + 2 {
		if lines[i] == "" {
			break
		}

		rover, err := model.CreateRoverFromLine(lines[i])
		if err != nil {
			log.Fatal(err)
		}

		rovers = append(rovers, rover)
		err = command.ApplyInstructions(plateau, rover, lines[i+1])
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, rover := range rovers {
		fmt.Printf("%d %d %s", rover.X, rover.Y, rover.Direction)
		fmt.Println("")
	}
}

func getLinesFromFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 3 {
		log.Fatal("Need at least 3 lines in the file")
	}

	return lines
}
