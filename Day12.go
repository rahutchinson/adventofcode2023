package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sumList(l []int) int {
	sum := 0

	// Iterate through the array and accumulate the sum
	for _, num := range l {
		sum += num
	}
	return sum
}

// ..#.### 1,1,3
func isValid(line string, arrange []int) bool {
	line = "." + line + "."
	currentIndex := 0
	count := strings.Count(line, "#")
	if count > sumList(arrange) || count < sumList(arrange) {
		return false
	}
	for i := 0; i < len(arrange); i++ {
		if currentIndex == len(line)-2 {
			return false
		}
		tofind := strings.Repeat("#", arrange[i])
		tofind = "." + tofind + "."
		index := strings.Index(line[currentIndex:], tofind)
		if index != -1 {
			count++
			currentIndex = index
		} else {
			return false
		}
	}
	return true
}

func generateCombinations(line string, minNumHash int) []string {
	var combinations []string
	seen := make(map[string]bool)

	var backtrack func(combination string, index, hashCount int)
	backtrack = func(combination string, index, hashCount int) {
		if index == len(line) {
			if hashCount >= minNumHash && !seen[combination] {
				combinations = append(combinations, combination)
				seen[combination] = true
			}
			return
		}

		if line[index] == '?' {
			backtrack(combination+".", index+1, hashCount)   // Replace '?' with '.'
			backtrack(combination+"#", index+1, hashCount+1) // Replace '?' with '#'
		} else {
			backtrack(combination+string(line[index]), index+1, hashCount)
		}
	}

	backtrack("", 0, 0)
	// Initialize a slice to store the keys
	var keys []string

	// Iterate over the map and collect the keys
	for key := range seen {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	dat, err := os.ReadFile("./inputs/day12/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	running := 0
	for _, line := range temp {
		splits := strings.Split(line, " ")
		vs := strings.Split(splits[1], ",")
		vsi := make([]int, 0, len(vs))
		for _, val := range vs {
			v, _ := strconv.Atoi(val)
			vsi = append(vsi, v)
		}
		l := splits[0]
		vArs := 0
		plines := generateCombinations(l, 1)
		for _, pline := range plines {
			// fmt.Printf("line: %s %t\n", pline, isValid(pline, vsi))
			if isValid(pline, vsi) {
				vArs++
			}
		}
		running += vArs
	}
	fmt.Println(running)
}

// 47935 too high
