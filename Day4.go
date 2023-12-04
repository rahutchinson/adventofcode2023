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

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func correctNumbers(correct, guess string) int {
	guesses := strings.Split(strings.TrimSpace(guess), " ")
	corrects := strings.Split(strings.TrimSpace(correct), " ")
	count := 0
	for _, g := range guesses {
		if contains(corrects, g) {
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

func main() {
	dat, err := os.ReadFile("./inputs/day4/example.txt")
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
