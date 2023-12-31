package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const input = "99,56,7,15,81,26,75,40,87,59,62,24,58,34,78,86,44,65,18,94,20,17,98,29,57,92,14,32,46,79,85,84,35,68,55,22,41,61,90,11,69,96,23,47,43,80,72,50,97,33,53,25,28,51,49,64,12,63,21,48,27,19,67,88,66,45,3,71,16,70,76,13,60,77,73,1,8,10,52,38,36,74,83,2,37,6,31,91,89,54,42,30,5,82,9,95,93,4,0,39"

func loadBoards() []*bingoBoard {
	file, err := os.Open("pkg/2021/day4/boards.txt")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var boards []*bingoBoard
	var buffer [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			split := strings.Split(line, " ")
			var partial []int
			for _, numStr := range split {
				if numStr != "" {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						panic(err)
					}

					partial = append(partial, num)
				}
			}

			buffer = append(buffer, partial)
		} else {
			// Next board
			boards = append(boards, newBingoBoard(buffer, len(boards)))
			buffer = [][]int{}
		}
	}

	return boards
}

func Part1() {
	start := time.Now()
	boards := loadBoards()

iteration:
	for _, numberStr := range strings.Split(input, ",") {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}

		fmt.Println("----------------------------------")
		fmt.Println("Marking boards with ballot", number)
		for _, board := range boards {
			if board.MarkNumber(number) {
				// We have a winning board
				board.PrintBoard()
				fmt.Println("Winning board:", board.BoardIndex, board.LastMarkedNumber()*board.SumUnmarked())
				break iteration
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func removeBoard(boards []*bingoBoard, boardIndex int) []*bingoBoard {
	bingoBoards := boards[:0]
	for _, board := range boards {
		if board.BoardIndex != boardIndex {
			bingoBoards = append(bingoBoards, board)
		}
	}
	return bingoBoards
}

func Part2() {
	start := time.Now()
	boards := loadBoards()

	var winnerBoards []*bingoBoard

	for _, numberStr := range strings.Split(input, ",") {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}

		fmt.Println("----------------------------------")
		fmt.Println("Marking boards with ballot", number)

		for _, board := range boards {
			if board.MarkNumber(number) {
				// We have a winning board
				winnerBoards = append(winnerBoards, board)
			}
		}

		// Remove the winning boards from the array
		for _, board := range winnerBoards {
			boards = removeBoard(boards, board.BoardIndex)
		}

		fmt.Println(len(boards), "boards remain")
		if len(boards) == 0 {
			break
		}
	}

	// Get the last winner from the array
	lastWinner := winnerBoards[len(winnerBoards)-1]
	lastWinner.PrintBoard()
	fmt.Println("Last winning board:", lastWinner.BoardIndex, lastWinner.LastMarkedNumber()*lastWinner.SumUnmarked())

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
