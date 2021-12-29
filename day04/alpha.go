package day04

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"advent21"
)

type gameBoard []int

var gameBoards []gameBoard

func Alpha() {
	file, err := advent21.OpenFile("day04/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer advent21.CloseFile(file)
	scanner := bufio.NewScanner(file)
	//winInCalls := 99
	lineCounter := 0
	boardCounter := 0
	currentBoard := make(gameBoard, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if lineCounter == 0 {
			fmt.Println("Call iterator")
			callIterator()
		} else {
			if line == "" {
				if len(currentBoard) > 0 {
					boardCounter++
					gameBoards = append(gameBoards, currentBoard)
					if len(currentBoard) != 25 {
						err = fmt.Errorf("current board should have 25 numbers but instead has %v", len(currentBoard))
						panic(err.Error())
					}
					fmt.Println("Board iterator", boardCounter)
					fmt.Println(currentBoard)
					currentBoard = make(gameBoard, 0)

				}
				continue
			}
			boardLine := boardIterator(line)
			currentBoard = append(currentBoard, boardLine...)
		}
		lineCounter++
	}
	fmt.Println("Number of boards:", len(gameBoards))
}

func callIterator() {}

func boardIterator(line string) (boardSlice gameBoard) {
	lineSlice := strings.Split(line, " ")
	for _, datum := range lineSlice {
		if datum == "" {
			continue
		}
		currentDatum, err := strconv.Atoi(datum)
		if err != nil {
			panic("failed iterating on a board: " + err.Error())
		}
		boardSlice = append(boardSlice, currentDatum)
	}

	return boardSlice
}

func (board gameBoard) winChecker(calls []int) {

}
