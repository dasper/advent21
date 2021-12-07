package day02

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"advent21"
)

func Alpha() {
	var position struct {
		horizontal int
		depth      int
	}
	file, err := advent21.OpenFile("day02/data.asc")
	if err != nil {
		log.Fatal(err)
	}
	defer advent21.CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		currentMeasurement, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
		}
		switch words[0] {
		case "forward":
			position.horizontal = position.horizontal + currentMeasurement
		case "down":
			position.depth = position.depth + currentMeasurement
		case "up":
			position.depth = position.depth - currentMeasurement
		}
	}

	fmt.Println("depth", position.depth)
	fmt.Println("horizontal", position.horizontal)
	fmt.Println("product", position.horizontal*position.depth)
}
