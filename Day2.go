package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 12 red cubes, 13 green cubes, and 14 blue cubes
var globalMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// only the sets seperated by ;
func isPossible(line string) bool {
	sets := strings.Split(line, ";")
	for _, valu := range sets {
		num_color := strings.Split(valu, ",")
		for _, valus := range num_color {
			single_num_col := strings.Split(strings.TrimSpace(valus), " ")
			num, err := strconv.Atoi(single_num_col[0])
			check(err)
			if globalMap[single_num_col[1]] < num {
				return false
			}
		}
	}
	return true
}

func main() {
	dat, err := os.ReadFile("./inputs/day2/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	running_sum := 0
	for _, val := range temp {
		split_string := strings.Split(val, ":")
		game_id, err := strconv.Atoi(strings.Split(split_string[0], " ")[1])
		check(err)
		if isPossible(split_string[1]) {
			running_sum += game_id
		}
	}
	fmt.Println(running_sum)
}
