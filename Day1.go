package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var globalNumMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findAllIndexes(fullString, substring string, num string) map[int]string {
	word_with_index := make(map[int]string)
	for offset := 0; ; {
		index := strings.Index(fullString[offset:], substring)
		if index == -1 {
			break
		}
		word_with_index[offset+index] = num
		offset += index + 1
	}
	return word_with_index
}

func replaceAtIndex(str string, index int, replacement string) string {
	if index < 0 || index >= len(str) {
		return str // Index out of range, return original string
	}
	return str[:index] + string(replacement) + str[index+1:]
}

func removeNonNumbersWithWordSearch(input string) string {
	full_map := make(map[int]string)
	for key, num := range globalNumMap {
		singleMap := findAllIndexes(input, key, num)
		for key, value := range singleMap {
			full_map[key] = value
		}
	}

	final_string := input
	for index, num := range full_map {
		final_string = replaceAtIndex(final_string, index, num)
	}

	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(final_string, "")
}

func main() {
	dat, err := os.ReadFile("./inputs/day1/input-A.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")

	running_sum := 0
	for _, value := range temp {
		only_nums := removeNonNumbersWithWordSearch(value)
		combinded, err := strconv.Atoi(string(only_nums[0]) + string(only_nums[len(only_nums)-1]))
		fmt.Println(combinded)
		check(err)
		running_sum += combinded

	}
	fmt.Println(running_sum)
}
