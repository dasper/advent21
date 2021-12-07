package day01

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"advent21"
)

func Alpha() {
	file, err := advent21.OpenFile("day01/data.asc")
	if err != nil {
		log.Fatal(err)
	}
	defer advent21.CloseFile(file)

	scanner := bufio.NewScanner(file)
	previousMeasurement := 0
	tally := 0
	started := false
	for scanner.Scan() {
		currentMeasurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if started && previousMeasurement < currentMeasurement {
			tally++
		}
		previousMeasurement = currentMeasurement
		started = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("larger measurements:", tally)
}
