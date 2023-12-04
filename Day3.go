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

func isSymbol(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}

func isNumber(r rune) bool {
	return unicode.IsNumber(r)
}

func isSymbolNotPunct(r rune) bool {
	return !unicode.IsNumber(r) && !(r == rune('.'))
}

func isStar(r rune) bool {
	return r == rune('*')
}

func splitOnSymbols(s string) []string {
	var parts []string
	var currentPart []rune

	for _, r := range s {
		if isSymbol(r) {
			if len(currentPart) > 0 {
				parts = append(parts, string(currentPart))
				currentPart = []rune{}
			}
		} else {
			currentPart = append(currentPart, r)
		}
	}

	if len(currentPart) > 0 {
		parts = append(parts, string(currentPart))
	}

	return parts
}

func numberFoundSymbol(index, length, row int, array []string) bool {
	maxRow := len(array) - 1
	maxCol := len(array[row]) - 1
	for i := 0; i < length; i++ {
		// all but last row
		if row < maxRow {
			// if row <= 140
			// check down
			if isSymbolNotPunct(rune(array[row+1][index+i])) {
				return true
			}
		}
		//Check up all but first
		if row > 0 {
			if isSymbolNotPunct(rune(array[row-1][index+i])) {
				return true
			}
		}

	}
	if index+length < maxCol {
		// check right
		if isSymbolNotPunct(rune(array[row][index+length])) {
			return true
		}
		//right Diags
		if row != maxRow && isSymbolNotPunct(rune(array[row+1][index+length])) {
			return true
		}
		if row != 0 && isSymbolNotPunct(rune(array[row-1][index+length])) {
			return true
		}
	}
	//Check left
	if index > 0 {
		if isSymbolNotPunct(rune(array[row][index-1])) {
			return true
		}
		//right Diags
		if row != maxRow && isSymbolNotPunct(rune(array[row+1][index-1])) {
			return true
		}
		if row != 0 && isSymbolNotPunct(rune(array[row-1][index-1])) {
			return true
		}
	}
	return false
}

func getFullNumber(index, row int, foundNums map[Pair]int) int {
	return foundNums[Pair{row, index}]
}

func removeDuplicates(s []int) []int {
	m := make(map[int]bool)
	for _, v := range s {
		m[v] = true
	}
	var result []int
	for k := range m {
		result = append(result, k)
	}
	return result
}

func symbolFoundNumber(index, row int, array []string, foundNums map[Pair]int) int {
	maxRow := len(array) - 1
	maxCol := len(array[row]) - 1
	// all but last row
	var found []int
	if row < maxRow {
		// check down
		if isNumber(rune(array[row+1][index])) {
			found = append(found, getFullNumber(index, row+1, foundNums))
		}
	}
	//Check up all but first
	if row > 0 {
		if isNumber(rune(array[row-1][index])) {
			found = append(found, getFullNumber(index, row-1, foundNums))
		}
	}

	if index < maxCol {
		// check right
		if isNumber(rune(array[row][index+1])) {
			found = append(found, getFullNumber(index+1, row, foundNums))
		}
		//right Diags
		if row != maxRow && isNumber(rune(array[row+1][index+1])) {
			found = append(found, getFullNumber(index+1, row+1, foundNums))
		}
		fmt.Println(isNumber(rune(array[row-1][index+1])))
		if row != 0 && isNumber(rune(array[row-1][index+1])) {
			found = append(found, getFullNumber(index+1, row-1, foundNums))
		}
	}
	//Check left
	if index > 0 {
		if isNumber(rune(array[row][index-1])) {
			found = append(found, getFullNumber(index-1, row, foundNums))
		}
		//left Diags
		if row != maxRow && isNumber(rune(array[row+1][index-1])) {
			found = append(found, getFullNumber(index-1, row+1, foundNums))
		}

		if row != 0 && isNumber(rune(array[row-1][index-1])) {
			found = append(found, getFullNumber(index-1, row-1, foundNums))
		}
	}
	found = removeDuplicates(found)
	if len(found) == 2 {
		fmt.Println(found)
		return found[0] * found[1]
	} else {
		return 0
	}
}

type Pair struct {
	Row   int
	Index int
}

func main() {
	dat, err := os.ReadFile("./inputs/day3/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")

	index_num_map := make(map[Pair]int)
	for row, line := range temp {
		vals := splitOnSymbols(line)
		running_index := 0
		for _, nums := range vals {
			num, err := strconv.Atoi(nums)
			isNumber := err == nil
			if isNumber {
				index := strings.Index(line[running_index:], nums) + running_index
				running_index = index + len(nums)
				for i := 0; i < len(nums); i++ {
					newPair := Pair{row, index + i}
					index_num_map[newPair] = num
				}
			}
		}
	}
	fmt.Println(index_num_map)

	running_sum := 0
	for row, line := range temp {
		lastIndex := 0
		fmt.Println("______________________" + string(row))
		for _, stars := range line {
			if isStar(rune(stars)) {
				index := strings.Index(line[lastIndex:], string(stars)) + lastIndex
				lastIndex = index + 1
				fmt.Println(index)
				running_sum += symbolFoundNumber(index, row, temp, index_num_map)
			}
		}

	}
	fmt.Println(running_sum)
}
