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

func checkZs(zss []string) bool {
	for _, val := range zss {
		if val[2] != 'Z' {
			return false
		}
	}
	return true
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

	var allA []string
	for key, _ := range mapping {
		if key[2] == 'A' {
			allA = append(allA, key)
		}
	}
	fmt.Println(allA)

	startTime := time.Now()
	steps := 0
	i := 0
	stepsMap := make(map[string]int)
	for _, val := range allA {
		allA_1 := []string{val}
		for !checkZs(allA_1) {
			allZs := make([]string, 0, 6)
			for _, a := range allA_1 {
				if pattern[i] == 'L' {
					allZs = append(allZs, mapping[a].l)
				} else {
					allZs = append(allZs, mapping[a].r)
				}
			}
			allA_1 = allZs
			steps++
			i = (i + 1) % 263

		}
		fmt.Println(steps)
		stepsMap[val] = steps
		duration := time.Since(startTime)
		fmt.Println("Execution time:", duration)
	}
	fmt.Println(stepsMap)
	duration := time.Since(startTime)
	fmt.Println("Execution time:", duration)
}
