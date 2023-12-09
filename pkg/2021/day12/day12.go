package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func loadInput() []string {
	file, err := os.Open("pkg/2021/day12/input.txt")
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
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func Part1() {
	start := time.Now()
	paths := loadInput()

	graph := NewGraph()

	for _, path := range paths {
		var source Node
		var dest Node

		split := strings.Split(path, "-")

		// lower is 1 visit only
		source = Node{
			value: split[0],
			visit: 1,
		}
		if isUpper(split[0]) {
			// Upper is unlimited visits
			source.visit = -1
		}
		// lower is 1 visit only
		dest = Node{
			value: split[1],
			visit: 1,
		}
		if isUpper(split[1]) {
			// Upper is unlimited visits
			dest.visit = -1
		}
		graph.AddEdge(source, dest)
		graph.AddEdge(dest, source)
	}

	graph.FindAllPaths(
		Node{value: "start", visit: 1},
		Node{value: "end", visit: 1},
	)

	fmt.Println("------------------------------------")
	fmt.Println("Count:", graph.PathCount())

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func Part2() {
	start := time.Now()
	paths := loadInput()

	var firstLower string
	for _, path := range paths {
		split := strings.Split(path, "-")
		if isLower(split[0]) && split[1] == "start" {
			firstLower = split[0]
			break
		}
		if isLower(split[1]) && split[0] == "start" {
			firstLower = split[1]
			break
		}
	}

	graph := NewGraph()

	for _, path := range paths {
		var source Node
		var dest Node

		split := strings.Split(path, "-")

		// lower is 1 visit only
		source = Node{
			value: split[0],
			visit: 1,
		}
		if isUpper(split[0]) {
			// Upper is unlimited visits
			source.visit = -1
		}
		if split[0] == firstLower {
			source.visit = 2
		}
		// lower is 1 visit only
		dest = Node{
			value: split[1],
			visit: 1,
		}
		if isUpper(split[1]) {
			// Upper is unlimited visits
			dest.visit = -1
		}
		if split[1] == firstLower {
			dest.visit = 2
		}

		graph.AddEdge(source, dest)
		graph.AddEdge(dest, source)
	}

	graph.FindAllPaths(
		Node{value: "start", visit: 1},
		Node{value: "end", visit: 1},
	)

	fmt.Println("------------------------------------")
	fmt.Println("Count:", graph.PathCount())

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
