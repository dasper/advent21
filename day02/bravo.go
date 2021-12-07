package day02

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"advent21"
)

func Bravo() {
	var position struct {
		aim        int
		horizontal int
		depth      int
	}
	file, err := advent21.OpenFile("day02/data.txt")
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
			position.depth = (position.aim * currentMeasurement) + position.depth
		case "down":
			position.aim = position.aim + currentMeasurement
		case "up":
			position.aim = position.aim - currentMeasurement
		}
	}

	fmt.Println("depth", position.depth)
	fmt.Println("horizontal", position.horizontal)
	fmt.Println("product", position.depth*position.horizontal)
}
