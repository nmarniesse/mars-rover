package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/nmarniesse/mars-rover/model"
)

func main() {
	filename := "./input/challenge.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	lines := strings.Split(string(content), "\n")
	if len(lines) < 2 {
		log.Fatal("Need at least 2 lines in the file")
	}

	firstLine := lines[0]
	model.CreatePlateauFromLine(firstLine)

	fmt.Println("Plateau created")

}
