package day03

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"advent21"
)

func Alpha() {
	file, err := advent21.OpenFile("day03/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer advent21.CloseFile(file)
	scanner := bufio.NewScanner(file)
	lines := 0
	tallies := make(map[int]int)
	gammaString := ""
	epsilonString := ""
	firstRunCompleted := false
	for scanner.Scan() {
		text := scanner.Text()
		for i, c := range text {
			if !firstRunCompleted {
				_, ok := tallies[i]
				if !ok {
					tallies[i] = 0
				}
			}
			if string(c) == "1" {
				tallies[i] = tallies[i] + 1
			}

		}
		lines++
		if !firstRunCompleted {
			firstRunCompleted = true
		}
	}

	for i := 0; i < len(tallies); i++ {
		v := tallies[i]
		if v*2 >= lines {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}
	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)
	fmt.Println("gamma", gamma)
	fmt.Println("epsilon", epsilon)
	fmt.Println("product", gamma*epsilon)
}
