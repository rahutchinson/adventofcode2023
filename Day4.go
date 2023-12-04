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

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func removeDuplicatesString(s []string) []string {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

func makeIntList(list []string) []int {
	var int_list []int
	for _, str := range list {
		val, err := strconv.Atoi(string(str))
		if err == nil {
			int_list = append(int_list, val)
		}
	}
	return int_list
}

func correctNumbers(winner, guess string) int {
	guesses := strings.Split(strings.TrimSpace(guess), " ")
	winners := strings.Split(strings.TrimSpace(winner), " ")
	winners_int := makeIntList(winners)
	guesses_int := makeIntList(guesses)
	count := 0
	for _, g := range winners_int {
		if contains(guesses_int, g) {
			count += 1
		}
	}
	if count > 0 {
		r_value := 1
		for i := 1; i < count; i++ {
			r_value *= 2
		}
		return r_value
	}
	return 0
}

// 59722 too high
// 41937 too high
func main() {
	dat, err := os.ReadFile("./inputs/day4/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")

	running_sum := 0
	for _, value := range temp {
		tClean := strings.Split(string(value), ":")[1]
		tempClean := strings.Split(tClean, "|")
		running_sum += correctNumbers(tempClean[0], tempClean[1])
	}
	fmt.Println(running_sum)
}
