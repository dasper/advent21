package day01

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"advent21"
)

func Bravo() {
	file, err := advent21.OpenFile("day01/data.asc")
	if err != nil {
		log.Fatal(err)
	}
	defer advent21.CloseFile(file)

	scanner := bufio.NewScanner(file)
	slidingLimit := 4
	slider := make([]int, 0)
	tally := 0

	for scanner.Scan() {
		currentMeasurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		slider = append(slider, currentMeasurement)
		if len(slider) == slidingLimit {
			previousSum := 0
			currentSum := 0
			for _, value := range slider[:len(slider)-1] {
				previousSum = previousSum + value
			}
			for _, value := range slider[1:] {
				currentSum = currentSum + value
			}
			if previousSum < currentSum {
				tally++
			}
			slider = slider[1:]
		}
	}

	fmt.Println("larger sums:", tally)
}
