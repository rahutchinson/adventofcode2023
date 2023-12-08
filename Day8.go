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
	dat, err := os.ReadFile("./inputs/day8/example.txt")
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

	steps := 0
	i := 0
	startTime := time.Now()
	for !checkZs(allA) {
		allZs := make([]string, 0, 6)
		for _, a := range allA {
			if pattern[i] == 'L' {
				allZs = append(allZs, mapping[a].l)
			} else {
				allZs = append(allZs, mapping[a].r)
			}
		}
		allA = allZs
		steps++
		i++
		if i == len(pattern) {
			i = 0
		}
	}
	duration := time.Since(startTime)
	fmt.Println("Execution time:", duration)
	fmt.Println(steps)
}
