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

func toInts(line []string) []int {
	ints := make([]int, 0, len(line))
	for i := 0; i < len(line); i++ {
		val := 0
		val, _ = strconv.Atoi(string(line[i]))
		ints = append(ints, val)
	}
	return ints
}

// Function to check if all values in a slice are zero
func allZero(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}
	return true
}

func findDiff(line []int) []int {
	diff := make([]int, 0, len(line))
	for i := 0; i < len(line)-1; i++ {
		diff = append(diff, line[i+1]-line[i])
	}
	return diff
}

func findNext(list []int) []int {
	if allZero(list) {
		return append([]int{0}, list...)
	} else {
		last := findNext(findDiff(list))
		fmt.Println(last)
		return append([]int{list[0] - last[0]}, list...)
	}
}

func main() {
	dat, err := os.ReadFile("./inputs/day9/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	running_sum := 0
	for _, line := range temp {
		fmt.Println("___________________________________________________________")
		final := findNext(toInts(strings.Split(line, " ")))
		fmt.Println(line)
		fmt.Println(final)
		running_sum += final[0]
	}
	fmt.Println(running_sum)
}

// 2105961943
// 2105961947 too high
//.2105961947
