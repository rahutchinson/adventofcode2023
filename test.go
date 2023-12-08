package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type KeyPair struct {
	l string
	r string
}

func main() {
	dat, err := os.ReadFile("./inputs/day8/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	pattern := temp[0]

	mapping := make(map[string]KeyPair)
	for _, line := range temp[2:] {
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, ",", "")
		line_clean := strings.Split(line, " = ")
		dirs := strings.Split(line_clean[1], " ")
		mapping[line_clean[0]] = KeyPair{l: dirs[0], r: dirs[1]}
	}

	allA := make([]string, 0)
	for key := range mapping {
		if key[2] == 'A' {
			allA = append(allA, key)
		}
	}

	startTime := time.Now()
	steps := 0
	patternLength := len(pattern)
	history := make(map[string]map[string]bool)

	for _, a := range allA {
		history[a] = make(map[string]bool)
	}

	loopDetected := false
	for !loopDetected {
		allZs := make([]string, len(allA))
		for index, a := range allA {
			next := mapping[a].l
			if pattern[steps%patternLength] == 'R' {
				next = mapping[a].r
			}

			if history[a] == nil {
				history[a] = make(map[string]bool)
			}

			if history[a][next] {
				fmt.Println("Loop detected for:", a)
				loopDetected = true
				break
			}

			history[a][next] = true
			allZs[index] = next
		}

		if loopDetected {
			break
		}

		allA = allZs
		steps++
	}

	duration := time.Since(startTime)
	fmt.Println("Steps:", steps)
	fmt.Println("Execution time:", duration)
}
