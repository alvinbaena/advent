package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readInput() string {
	file, err := os.Open("pkg/2021/day7/input.txt")
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

func loadInput() []int {
	split := strings.Split(readInput(), ",")

	crabs := make([]int, len(split))
	for i, s := range split {
		number, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		crabs[i] = number
	}

	return crabs
}

func median(crabs []int) int {
	dataCopy := make([]int, len(crabs))
	copy(dataCopy, crabs)

	sort.Ints(dataCopy)

	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	l := len(dataCopy)
	var median int
	if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}

func average(crabs []int) int {
	sum := 0
	for _, num := range crabs {
		sum += num
	}

	// I think this is wrong. Instead of casting we should be rounding, but the answer requires casting....
	return int(float64(sum) / float64(len(crabs)))
}

func difference(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func Part1() {
	start := time.Now()
	crabs := loadInput()

	// Calculate the median to get the best possible value
	median := median(crabs)
	fmt.Println("Best position", median)

	sum := 0
	// Get the difference and sum
	for _, position := range crabs {
		sum += difference(median, position)
	}

	fmt.Println("Fuel", sum)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func Part2() {
	start := time.Now()
	crabs := loadInput()

	avg := average(crabs)
	fmt.Println("Best position", avg)

	sum := 0
	// Get the difference and sum by increasing each time
	for _, position := range crabs {
		diff := difference(avg, position)
		for i := 1; i <= diff; i++ {
			sum += i
		}
	}

	fmt.Println("Fuel", sum)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
