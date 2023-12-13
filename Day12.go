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
			fmt.Printf("%s in %s at %d\n", tofind, line, index)
		} else {
			return false
		}
	}
	return true
}

func generateCombinations(line string, add bool, max int) []string {
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
	if add {
		for key := range seen {
			keys = append(keys, strings.Repeat(key+"?", 5))
		}
	} else {
		for key := range seen {
			keys = append(keys, key)
		}
	}
	return keys
}

func newArr(l []int) []int {
	repeatedArray := []int{}

	// Repeat the input array 5 times in order
	for i := 0; i < 5; i++ {
		repeatedArray = append(repeatedArray, l...)
	}
	return repeatedArray
}

func main() {
	// fmt.Println(isValid("...#.######..#####.", []int{1, 6, 5}))
	// return

	dat, err := os.ReadFile("./inputs/day12/example.txt")
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
		plines := generateCombinations(l, true, sumList(vsi))
		fmt.Println(len(plines))
		vsi2 := newArr(vsi)
		// fmt.Println(vsi2)
		for _, pline := range plines {
			plines2 := generateCombinations(pline, false, sumList(vsi2))
			fmt.Println(len(plines2))
			for _, p := range plines2 {
				// fmt.Printf("line: %s %t\n", p, isValid(p, vsi2))
				if isValid(p, vsi2) {
					// fmt.Printf("line: %s %t\n", pline, isValid(pline, vsi))
					vArs++
				}
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
