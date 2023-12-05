package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Map struct {
	source      int
	destination int
	length      int
}

func trimAtoi(toTrim string) int {
	num, _ := strconv.Atoi(strings.TrimSpace(toTrim))
	return num

}

func main() {
	dat, err := os.ReadFile("./inputs/day5/example.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	var seeds []int
	mapOfLists := make(map[string][]Map)
	currentMap := ""
	for _, line := range temp {
		split := strings.Split(line, ":")
		if split[0] == "seeds" {
			nums := strings.Split(strings.TrimSpace(split[1]), " ")
			for _, num := range nums {
				conv, _ := strconv.Atoi(num)
				seeds = append(seeds, conv)
			}
			continue
		}
		if line == "" {
			continue
		}
		if unicode.IsLetter(rune(line[0])) {
			currentMap = split[0]
			continue
		}
		if unicode.IsNumber(rune(line[0])) {
			nums := strings.Split(strings.TrimSpace(line), " ")
			mapOfLists[currentMap] = append(mapOfLists[currentMap], Map{trimAtoi(nums[0]), trimAtoi(nums[1]), trimAtoi(nums[2])})

		}
	}
	fmt.Println(seeds)
	fmt.Println(mapOfLists)

}
