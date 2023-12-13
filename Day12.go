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
	if count != sumList(arrange) {
		return false
	}
	for i := 0; i < len(arrange); i++ {
		line = line[currentIndex:]
		tofind := strings.Repeat("#", arrange[i])
		tofind = "." + tofind + "."
		index := strings.Index(line, tofind)
		if index != -1 {
			count++
			currentIndex = index + len(tofind) - 1
			// fmt.Printf("%s in %s at %d\n", tofind, line, index)
		} else {
			return false
		}
	}
	return true
}

func generateCombinations(line string) []string {
	seen := make(map[string]bool)

	var backtrack func(combination string, index int)
	backtrack = func(combination string, index int) {
		if index == len(line) {
			if !seen[combination] {
				seen[combination] = true
			}
			return
		}

		if line[index] == '?' {
			backtrack(combination+".", index+1) // Replace '?' with '.'
			backtrack(combination+"#", index+1) // Replace '?' with '#'
		} else {
			backtrack(combination+string(line[index]), index+1)
		}
	}

	backtrack("", 0)
	// Initialize a slice to store the keys
	var keys []string

	// Iterate over the map and collect the keys
	for key := range seen {
		
		keys = append(keys, key)
	}
	return keys
}

func main() {
	// fmt.Println(isValid("...#.######..#####.", []int{1, 6, 5}))
	// return

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
		plines := generateCombinations(l)
		for _, pline := range plines {
			// fmt.Printf("line: %s %t\n", pline, isValid(pline, vsi))
			if isValid(pline, vsi) {
				// fmt.Printf("line: %s %t\n", pline, isValid(pline, vsi))
				vArs++
			}
		}
		running += vArs
		fmt.Println(vArs)
	}
	fmt.Println(running)
}

// 47935 too high
// 28442 too high
// 5890 to low
