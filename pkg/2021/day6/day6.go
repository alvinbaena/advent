package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput() string {
	file, err := os.Open("pkg/2021/day6/input.txt")
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
	var res string
	for scanner.Scan() {
		res = scanner.Text()
		if err != nil {
			panic(err)
		}
	}

	return res
}

func copulation(school []string, days int) uint64 {
	sea := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	// Fill the sea
	for _, f := range school {
		timer, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		sea[timer] = sea[timer] + 1
	}

	// Start the counter
	for i := 0; i < days; i++ {
		// After the day all 1's become 0's and reset.
		// Move positions
		newFishes := 0
		for timer, amount := range sea {
			if timer == 0 {
				// Become a 6 later
				newFishes = amount
			} else {
				sea[timer-1] = amount
			}
		}

		// Add more 6's and 8's
		sea[6] = sea[6] + newFishes
		sea[8] = newFishes
	}

	fmt.Println("Fishes", sea)

	count := uint64(0)
	for _, fishes := range sea {
		count += uint64(fishes)
	}

	return count
}

func Part1() {
	start := time.Now()
	school := strings.Split(readInput(), ",")
	fmt.Println("Count", copulation(school, 80))

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func Part2() {
	start := time.Now()
	elapsed := time.Since(start)

	school := strings.Split(readInput(), ",")
	fmt.Println("Count", copulation(school, 256))

	fmt.Println("Function call took ", elapsed)
}
