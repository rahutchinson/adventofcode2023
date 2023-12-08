package main

import (
	"fmt"
	"os"
	"strings"
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

	fmt.Println(len(mapping))
	fmt.Println(mapping["FLR"])
	loc := "AAA"
	endLoc := "ZZZ"
	steps := 0
	i := 0
	for i < len(pattern) {
		if loc == endLoc {
			fmt.Println(steps)
			break
		} else {
			if pattern[i] == 'L' {
				loc = mapping[loc].l
			} else {
				loc = mapping[loc].r
			}
			steps++
			i++
			if i == len(pattern) {
				i = 0
			}
		}
	}
	fmt.Println(steps)
}
