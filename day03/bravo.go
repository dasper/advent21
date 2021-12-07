package day03

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strconv"

	"advent21"
)

func Bravo() {
	file, err := advent21.OpenFile("day03/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer advent21.CloseFile(file)
	scanner := bufio.NewScanner(file)

	list := make([]string, 0)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	oxygenString, err := filterList(list, true)
	if err != nil {
		log.Fatal("Expected O2 match but got an error:", err.Error())
	}
	carbonDioxideString, _ := filterList(list, false)
	if err != nil {
		log.Fatal("Expected CO2 match but got an error:", err.Error())
	}

	oxygen, _ := strconv.ParseInt(oxygenString, 2, 64)
	carbonDioxide, _ := strconv.ParseInt(carbonDioxideString, 2, 64)
	fmt.Println("oxygen", oxygen)
	fmt.Println("carbonDioxide", carbonDioxide)
	fmt.Println("product", oxygen*carbonDioxide)
}

func filterList(list []string, tiebreaker bool) (response string, err error) {
	i := 0
	for {
		list = findMatchingBit(list, i, tiebreaker)
		if len(list) < 2 {
			break
		}
		i++
	}
	if len(list) == 0 {
		err = errors.New("failed to find match")
	}
	response = list[0]
	return
}

func findMatchingBit(list []string, position int, match bool) []string {
	var trueList, falseList []string
	for _, item := range list {
		if rune(item[position]) == '1' {
			trueList = append(trueList, item)
		} else {
			falseList = append(falseList, item)
		}
	}
	if len(trueList) >= len(falseList) == match {
		return trueList
	} else {
		return falseList
	}
}
