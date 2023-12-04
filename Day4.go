package main

import (
	"fmt"
	"os"
	"sort"
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

func correctNumbersNoMultiply(winner, guess string) int {
	guesses := strings.Split(strings.TrimSpace(guess), " ")
	winners := strings.Split(strings.TrimSpace(winner), " ")
	winners_int := makeIntList(winners)
	guesses_int := makeIntList(guesses)
	count := 1
	for _, g := range winners_int {
		if contains(guesses_int, g) {
			count += 1
		}
	}
	return count
}

func initializeCountMap(length int) map[int]int {
	count_map := make(map[int]int)
	for i := 1; i < length+1; i++ {
		count_map[i] = 1
	}
	return count_map
}

func processCardMapWinners(w_map map[int]int) int {
	fmt.Println(w_map)
	fmt.Println("__________________")
	count_map := initializeCountMap(len(w_map))
	keys := make([]int, 0, len(w_map))
	for key := range w_map {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// How many wins
	for _, key := range keys {
		for j := 0; j < count_map[key]; j++ {
			if w_map[key] != 1 || key == 1 {
				// adds the wins to the following cards
				for i := 1; i < w_map[key]; i++ {
					if !(key+i > len(w_map)) {
						count_map[key+i] = count_map[key+i] + 1
					}
				}
			}
		}
	}
	sum := 0
	for _, value := range count_map {
		sum += value
	}
	fmt.Println(count_map)
	return sum
}

// 10642 too low
func main() {
	dat, err := os.ReadFile("./inputs/day4/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")

	// running_sum := 0
	card_map := make(map[int]int)
	for row, value := range temp {
		tClean := strings.Split(string(value), ":")[1]
		tempClean := strings.Split(tClean, "|")
		card_map[row+1] = correctNumbersNoMultiply(tempClean[0], tempClean[1])
	}
	fmt.Println(processCardMapWinners(card_map))

}
