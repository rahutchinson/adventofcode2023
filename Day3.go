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

func isSymbolNotPunct(r rune) bool {
	return !unicode.IsNumber(r) && !(r == rune('.'))
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

func main() {
	dat, err := os.ReadFile("./inputs/day3/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")

	running_sum := 0
	for row, line := range temp {
		vals := splitOnSymbols(line)
		fmt.Print(row)
		fmt.Println(vals)
		lastIndex := 0
		for _, nums := range vals {
			num, err := strconv.Atoi(nums)
			isNumber := err == nil
			if isNumber {
				index := strings.Index(line[lastIndex+1:], nums) + lastIndex + 1
				lastIndex = index
				if numberFoundSymbol(index, len(nums), row, temp) {
					running_sum += num
				}
			}
		}

	}
	fmt.Println(running_sum)
}
